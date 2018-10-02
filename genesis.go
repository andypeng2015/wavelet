package wavelet

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/perlin-network/wavelet/log"
	"github.com/pkg/errors"
)

// LoadGenesisTransaction loads the genesis transaction from a json file.
func LoadGenesisTransaction(path string) ([]*Account, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var jsonEntries []map[string]interface{}
	if err := json.Unmarshal(byteValue, &jsonEntries); err != nil {
		return nil, err
	}

	var accounts []*Account
	for i, entry := range jsonEntries {
		encoded, ok := entry["public_key"]
		if !ok {
			return nil, errors.Errorf("Genesis file malformed, failed to find public_key for entry %d", i)
		}
		encodedID, ok := encoded.(string)
		if !ok {
			return nil, errors.Errorf("Genesis file malformed, failed to cast public_key for entry %d", i)
		}

		id, err := hex.DecodeString(encodedID)
		if err != nil {
			return nil, err
		}

		account := NewAccount(id)
		for key, v := range entry {
			if key == "public_key" {
				// we already processed this special entry, skip it
				continue
			}
			switch value := v.(type) {
			case uint64:
				account.Store(key, writeUint64(value))
			case string:
				account.Store(key, writeBytes(value))
			}
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

// ApplyGenesisTransactions applies accounts to the ledger
func ApplyGenesisTransactions(ledger *Ledger, accounts []*Account) {
	for i, account := range accounts {
		if err := ledger.SaveAccount(account, nil); err != nil {
			log.Fatal().Err(err).
				Str("public_key", string(account.PublicKey)).
				Int("index", i).
				Msg("Failed to save genesis account information.")
		}
	}
}
