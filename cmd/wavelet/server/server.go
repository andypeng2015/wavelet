package server

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/cipher"
	"github.com/perlin-network/noise/edwards25519"
	"github.com/perlin-network/noise/handshake"
	"github.com/perlin-network/noise/nat"
	"github.com/perlin-network/noise/skademlia"
	"github.com/perlin-network/wavelet"
	"github.com/perlin-network/wavelet/api"
	"github.com/perlin-network/wavelet/internal/snappy"
	"github.com/perlin-network/wavelet/log"
	"github.com/perlin-network/wavelet/store"
	"github.com/perlin-network/wavelet/sys"
	"github.com/perlin-network/wavelet/wctl"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Config struct {
	NAT      bool
	Host     string
	Port     uint
	Wallet   string // hex encoded
	Genesis  *string
	APIPort  uint
	Peers    []string
	Database string
}

var DefaultConfig = Config{
	Host:     "127.0.0.1",
	Port:     3000,
	Wallet:   "87a6813c3b4cf534b6ae82db9b1409fa7dbd5c13dba5858970b56084c4a930eb400056ee68a7cc2695222df05ea76875bc27ec6e61e8e62317c336157019c405",
	Genesis:  nil,
	APIPort:  9000,
	Peers:    []string{},
	Database: "",
}

type Wavelet struct {
	Logger zerolog.Logger
	Keys   *skademlia.Keypair
}

func Start(cfg *Config) error {
	if cfg == nil {
		cfg = &DefaultConfig
	}

	w := Wavelet{}

	// Make a logger
	w.Logger = log.Node()

	// TODO(diamond): change all panics to useful logger.Fatals

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("Failed to open port %d: %v", cfg.Port, err)
	}

	addr := net.JoinHostPort(
		cfg.Host, strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
	)

	if cfg.NAT {
		if len(cfg.Peers) > 1 {
			resolver := nat.NewPMP()

			if err := resolver.AddMapping("tcp",
				uint16(listener.Addr().(*net.TCPAddr).Port),
				uint16(listener.Addr().(*net.TCPAddr).Port),
				30*time.Minute,
			); err != nil {
				return fmt.Errorf("Failed to initialize NAT: %v", err)
			}
		}

		resp, err := http.Get("http://myexternalip.com/raw")
		if err != nil {
			return fmt.Errorf("Failed to get external IP: %v", err)
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to get external IP: %v", err)
		}

		addr = net.JoinHostPort(
			string(ip), strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		)
	}

	w.Logger.Info().Str("addr", addr).
		Msg("Listening for peers.")

	// Load keys
	var privateKey edwards25519.PrivateKey

	i, err := hex.Decode(privateKey[:], []byte(cfg.Wallet))
	if err != nil {
		return fmt.Errorf("Failed to decode hex wallet %s", cfg.Wallet)
	}

	if i != edwards25519.SizePrivateKey {
		return fmt.Errorf("Wallet is not of the right length (%d not %d)",
			i, edwards25519.SizePrivateKey)
	}

	keys, err := skademlia.LoadKeys(privateKey, sys.SKademliaC1, sys.SKademliaC2)
	if err != nil {
		return fmt.Errorf("The wallet specified is invalid: %v", err)
	}

	client := skademlia.NewClient(
		addr, keys,
		skademlia.WithC1(sys.SKademliaC1),
		skademlia.WithC2(sys.SKademliaC2),
		skademlia.WithDialOptions(grpc.WithDefaultCallOptions(
			grpc.UseCompressor(snappy.Name))),
	)

	client.SetCredentials(noise.NewCredentials(
		addr, handshake.NewECDH(), cipher.NewAEAD(), client.Protocol(),
	))

	client.OnPeerJoin(func(conn *grpc.ClientConn, id *skademlia.ID) {
		publicKey := id.PublicKey()

		logger := log.Network("joined")
		logger.Info().
			Hex("public_key", publicKey[:]).
			Str("address", id.Address()).
			Msg("Peer has joined.")

	})

	client.OnPeerLeave(func(conn *grpc.ClientConn, id *skademlia.ID) {
		publicKey := id.PublicKey()

		logger := log.Network("left")
		logger.Info().
			Hex("public_key", publicKey[:]).
			Str("address", id.Address()).
			Msg("Peer has left.")
	})

	kv, err := store.NewLevelDB(cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create/open database located at %q.", cfg.Database)
	}

	ledger := wavelet.NewLedger(kv, client, wavelet.WithGenesis(cfg.Genesis))

	go func() {
		server := client.Listen()

		wavelet.RegisterWaveletServer(server, ledger.Protocol())

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	for _, addr := range cfg.Peers {
		if _, err := client.Dial(addr); err != nil {
			logger.Warn().Err(err).
				Str("addr", addr).
				Msg("Error dialing")
		}
	}

	if peers := client.Bootstrap(); len(peers) > 0 {
		var ids []string

		for _, id := range peers {
			ids = append(ids, id.String())
		}

		logger.Info().Msgf("Bootstrapped with peers: %+v", ids)
	}

	if cfg.APIPort == 0 {
		cfg.APIPort = 9000
	}

	go api.New().StartHTTP(int(cfg.APIPort), client, ledger, keys)

	c, err := wctl.NewClient(wctl.Config{
		APIHost:    "localhost",
		APIPort:    uint16(cfg.APIPort),
		PrivateKey: keys.PrivateKey(),
	})

	if err != nil {
		logger.Fatal().Err(err).
			Uint("port", cfg.APIPort).
			Msg("Failed to connect to API")
	}

	shell, err := NewCLI(c)
	if err != nil {
		logger.Fatal().Err(err).
			Msg("Failed to spawn the CLI")
	}

	shell.Start()

}