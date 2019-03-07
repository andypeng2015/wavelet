package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/perlin-network/noise/identity/ed25519"
	"github.com/perlin-network/wavelet/wctl"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

var (
	LoadedWallet      = "Loaded wallet."
	GeneratedWallet   = "Generated a wallet."
	StartedAPI        = "Started HTTP API server."
	ListeningForPeers = "Listening for peers."
)

var _ io.Writer = (*node)(nil)

type node struct {
	keys   *ed25519.Keypair
	client *wctl.Client

	args              []string
	nodeHost          string
	nodePort, apiPort uint16
}

func (n *node) Write(buf []byte) (num int, err error) {
	var fields map[string]interface{}

	decoder := json.NewDecoder(bytes.NewReader(buf))
	decoder.UseNumber()

	err = decoder.Decode(&fields)
	if err != nil {
		return num, errors.Wrapf(err, "cannot decode field: %q", err)
	}

	if msg, exists := fields["message"]; exists {
		if msg, ok := msg.(string); ok {
			err = n.parseMessage(fields, msg)

			if err != nil {
				fmt.Printf("%v\n", err)
				return num, errors.Wrap(err, "failed to parse message")
			}
		}
	}

	return len(buf), nil
}

func (n *node) parseMessage(fields map[string]interface{}, msg string) error {
	switch msg {
	case GeneratedWallet:
		fallthrough
	case LoadedWallet:
		privateKey, err := hex.DecodeString(fields["privateKey"].(string))
		if err != nil {
			return errors.Wrap(err, "failed to decode nodes private key")
		}

		n.keys = ed25519.LoadKeys(privateKey)
	case StartedAPI:
		if n.keys == nil {
			return errors.New("started api before reading wallet keys")
		}

		if err := n.init(); err != nil {
			return errors.Wrap(err, "failed to init wavelet node")
		}
	}

	return nil
}

func (n *node) init() error {
	config := wctl.Config{
		APIHost: "127.0.0.1",
		APIPort: n.apiPort,
	}
	copy(config.RawPrivateKey[:], n.keys.PrivateKey())

	var err error

	if n.client, err = wctl.NewClient(config); err != nil {
		return errors.Wrap(err, "failed to init node HTTP API client")
	}

	if err = n.client.Init(); err != nil {
		return errors.Wrap(err, "failed to init session with HTTP API")
	}

	log.Info().
		Uint16("node_port", n.nodePort).
		Uint16("api_port", n.apiPort).
		Hex("public_key", n.keys.PublicKey()).
		Msg("Spawned a new Wavelet node.")

	return nil
}

func spawn(nodePort, apiPort uint16, randomWallet bool, peers ...string) *node {
	cmd := exec.Command("./wavelet", "-p", strconv.Itoa(int(nodePort)))

	if apiPort != 0 {
		cmd.Args = append(cmd.Args, "-api", strconv.Itoa(int(apiPort)))
	}

	if randomWallet {
		cmd.Args = append(cmd.Args, "-w", " ")
	}

	if len(peers) > 0 {
		cmd.Args = append(cmd.Args, strings.Join(peers, " "))
	}

	// TODO(kenta): allow external hosts
	n := &node{args: cmd.Args, nodeHost: "127.0.0.1", nodePort: nodePort, apiPort: apiPort}

	cmd.Stdout = n
	cmd.Stderr = n

	if err := cmd.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to spawn a single Wavelet node.")
	}

	return n
}
