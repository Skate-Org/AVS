package register

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/Skate-Org/AVS/contracts/bindings/IAVSDirectory"
	bindingIDelegationManager "github.com/Skate-Org/AVS/contracts/bindings/IDelegationManager"
	"github.com/Skate-Org/AVS/contracts/bindings/ISkateAVS"
	"github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/crypto/ecdsa"
	"github.com/Skate-Org/AVS/lib/logging"

	"github.com/Skate-Org/AVS/lib/on-chain/backend"
)

var logger = logging.NewLoggerWithConsoleWriter()

func RegisterOperatorWithAVS(
	privateKey *ecdsa.PrivateKey,
	config *cmd.EnvironmentConfig,
	salt [32]byte, expiry *big.Int,
) error {
	opAddr := ecdsa.PubkeyToAddress(privateKey.PublicKey)
	be, err := backend.NewBackend(config.HttpRPC)
	if err != nil {
		logger.Fatal("Can't instantiate backend", "rpcURL", config.HttpRPC)
		return err
	}

	avsDir, err := bindingIAVSDirectory.NewBindingIAVSDirectory(
		common.HexToAddress(config.AVSDirectory),
		be,
	)
	if err != nil {
		logger.Error("Can't bind avs directory", "address", config.AVSDirectory)
		return err
	}

	digestHash, err := avsDir.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{},
		opAddr,
		common.HexToAddress(config.SkateAVS),
		salt, expiry,
	)
	if err != nil {
		logger.Error("Failed to calculate AVS Registration hash", "error", err)
		return err
	}

	signature, err := ecdsa.Sign(digestHash[:], privateKey)
	if err != nil {
		return err
	}
	operatorSigantureWithSaltAndExpiry := bindingISkateAVS.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature[:],
		Salt:      salt,
		Expiry:    expiry,
	}

	avsContract, err := bindingISkateAVS.NewBindingISkateAVS(common.HexToAddress(config.SkateAVS), be)
	if err != nil {
		logger.Error("Can't bind avs contract", "address", config.SkateAVS)
		return err
	}

	ok, err := avsContract.CanRegister(
		&bind.CallOpts{}, ecdsa.PubkeyToAddress(privateKey.PublicKey),
	)
	if err != nil {
		logger.Fatal("Registration eligiblity", "allowed", ok, "error", err)
	}

	chainId := new(big.Int).SetUint64(config.MainChainId)

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	operatorAddress := ecdsa.PubkeyToAddress(privateKey.PublicKey)
	// Dry run transaction to check for potential rejections
	transactorNoSend := *transactor
	transactorNoSend.NoSend = true
	_, err = avsContract.RegisterOperatorToAVS(
		&transactorNoSend,
		operatorAddress,
		operatorSigantureWithSaltAndExpiry,
	)
	if err != nil {
		logger.Error("Transaction simulation failed", "error", errors.Wrap(err, "avsContract.RegisterOperator"))
		return err
	}

	tx, err := avsContract.RegisterOperatorToAVS(
		transactor, operatorAddress,
		operatorSigantureWithSaltAndExpiry,
	)
	if err != nil {
		logger.Error("Transaction failed", "error", errors.Wrap(err, "avsContract.RegisterOperator"))
		return err
	}
	account := ecdsa.PubkeyToAddress(privateKey.PublicKey).Hex()
	logger.Info("Registration request to Skate AVS sent...", "txHash", tx.Hash().Hex(), "account", account)

	receipt, err := backend.WaitMined(context.Background(), &be, tx)
	if err != nil {
		logger.Error("Failed to get transaction receipt", "error", errors.Wrap(err, "backend.TransactionReceipt"))
		return err
	}
	logger.Info("Registered to Skate AVS!", "status", receipt.Status, "gasUsed", receipt.GasUsed, "account", account)

	return nil
}

func RegisterOperatorWithEigenLayer(
	privateKey *ecdsa.PrivateKey,
	config *cmd.EnvironmentConfig,
	operatorDetails bindingIDelegationManager.IDelegationManagerOperatorDetails,
) error {
	be, err := backend.NewBackend(config.HttpRPC)
	if err != nil {
		return err
	}

	delegationManagerAddr := common.HexToAddress(config.DelegationManager)
	delegationManager, err := bindingIDelegationManager.NewBindingIDelegationManager(delegationManagerAddr, be)
	if err != nil {
		return err
	}

	chainId := new(big.Int).SetUint64(config.MainChainId)
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	// TODO: (futures version, if necessary) specify and upload metadata to IPFS.
	metadataURI := ""

	transactorNoSend := *transactor
	transactorNoSend.NoSend = true
	// Dry run transaction to check for potential rejections
	_, err = delegationManager.RegisterAsOperator(
		&transactorNoSend,
		operatorDetails,
		metadataURI,
	)
	if err != nil {
		logger.Error("Transaction simulation failed", "error", errors.Wrap(err, "delegationManager.RegisterAsOperator"))
		return err
	}

	tx, err := delegationManager.RegisterAsOperator(
		transactor,
		operatorDetails,
		metadataURI,
	)
	if err != nil {
		logger.Error("Transaction failed", "error", errors.Wrap(err, "delegationManager.RegisterAsOperator"))
		return err
	}

	account := ecdsa.PubkeyToAddress(privateKey.PublicKey).Hex()
	logger.Info("Registration request to EigenLayer sent...", "txHash", tx.Hash().Hex(), "account", account)
	receipt, err := backend.WaitMined(context.Background(), &be, tx)
	if err != nil {
		logger.Error("Failed to get transaction receipt", "error", errors.Wrap(err, "backend.TransactionReceipt"))
		return err
	}
	logger.Info("Registered to EigenLayer!", "status", receipt.Status, "gasUsed", receipt.GasUsed, "account", account)

	return nil
}
