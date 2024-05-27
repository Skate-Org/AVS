package bn254

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Similar structure for saving keys as ethereum keystore
// Id is omitted, retrieval by Pubkey instead
type encryptedBLSKeyJSONV3 struct {
	PubKey  string              `json:"pubKey"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
	Version int                 `json:"version"`
}

// SaveToFile saves the private key in an encrypted keystore file
func (k *BLSKey) SaveToFile(path string, password string) error {
	sk32Bytes := k.PrivKey.Bytes()
	cryptoStruct, err := keystore.EncryptDataV3(
		sk32Bytes[:],
		[]byte(password),
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	if err != nil {
		return err
	}

	encryptedBLSStruct := encryptedBLSKeyJSONV3{
		PubKey:  k.PubKey.String(),
		Crypto:  cryptoStruct,
		Version: 3,
	}
	data, err := json.Marshal(encryptedBLSStruct)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Println("Error creating directories:", err)
		return err
	}
	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func ReadPrivateKeyFromFile(path string, password string) (*BLSKey, error) {
	keyStoreContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	encryptedBLSStruct := &encryptedBLSKeyJSONV3{}
	err = json.Unmarshal(keyStoreContents, encryptedBLSStruct)
	if err != nil {
		return nil, err
	}

	// Check if pubkey is present, if not return error
	// There is an issue where if you specify ecdsa key file
	// it still works and returns a keypair since the format of storage is same.
	// This is to prevent and make sure pubkey is present.
	// ecdsa keys doesn't have that field
	if encryptedBLSStruct.PubKey == "" {
		return nil, fmt.Errorf("invalid bls key file. pubkey field not found")
	}

	skBytes, err := keystore.DecryptDataV3(encryptedBLSStruct.Crypto, password)
	if err != nil {
		return nil, err
	}

	privKey := new(fr.Element).SetBytes(skBytes)
	keyPair := NewKeyPair(privKey)
	return keyPair, nil
}
