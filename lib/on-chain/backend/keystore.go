package backend

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
)

const KeyStore = "keystore"

func DumpECDSAPrivateKeyToKeystore(privateKeyHex string, passphrase string) common.Address {
	logger := logging.NewLoggerWithConsoleWriter()
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logger.Fatal("Invalid private key", "err", errors.Wrap(err, "crpyo.HexToECDSA"))
	}

	// Create a new keystore directory
	ks := keystore.NewKeyStore(KeyStore, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import the private key into the keystore, encrypting it with the passphrase
	account, err := ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		logger.Fatal("Import ECDSA key failed", "err", errors.Wrap(err, "keystore.ImportECDSA"))
	}

	logger.Info("Keystore created for account:", "account", account.Address.Hex())
  return account.Address
}

func TransactorFromKeystore(address common.Address, passphrase string, chainId *big.Int) (*bind.TransactOpts, error) {
	ks := keystore.NewKeyStore(KeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	account := accounts.Account{Address: address}
	// keyJSON, err := ks.Export(account, passphrase, passphrase)
	ks.TimedUnlock(account, passphrase, time.Minute*0)
	return bind.NewKeyStoreTransactorWithChainID(ks, account, chainId)
}

func PrivateKeyFromKeystore(address common.Address, passphrase string) (*ecdsa.PrivateKey, error) {
	// Load the keystore
	ks := keystore.NewKeyStore(KeyStore, keystore.StandardScryptN, keystore.StandardScryptP)

	// Get the account corresponding to the given address
	account := accounts.Account{Address: address}

	// Export the key from the keystore
	keyJSON, err := ks.Export(account, passphrase, passphrase)
	if err != nil {
		return nil, errors.Wrap(err, "ks.Export")
	}

	// Decrypt the key using the passphrase
	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		return nil, errors.Wrap(err, "ks.Decrypt")
	}

	return key.PrivateKey, nil
}
