package cmd

// NOTE: command for operator to register with Skate AVS
import (
	"crypto/rand"
	"math/big"
	"time"

	bindingIDelegationManager "github.com/Skate-Org/AVS/contracts/bindings/IDelegationManager"
	libHash "github.com/Skate-Org/AVS/lib/crypto/hash"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/operator/register"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
)

func registerAvsCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string

	cmd := &cobra.Command{
		Use:   "register-avs",
		Short: "Register an operator with Skate AVS",
		Long: `Register an operator to Skate AVS. Note that the operator must already be registered with Eigen-Layer
    and there must be AT LEAST SOME STAKED AMOUNT.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}

			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/operator", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}

			if signerConfig.Address == "" {
				logger.Fatal("No signer provided, run with read-only mode")
			}

			privateKey, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
			if err != nil {
				logger.Fatal("Invalid keystore for signer", signerConfig)
				return err
			}

			b := make([]byte, 128)
			_, err = rand.Read(b)
			salt := libHash.Keccak256([]byte("very random string"), b)

			// Expires in 1 minutes (5 blocks)
			expiry := new(big.Int).SetInt64(time.Now().Unix() + 60)

			register.RegisterOperatorWithAVS(
				privateKey,
				envConfig,
				[32]byte(salt),
				expiry,
			)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)

	return cmd
}

func registerEigenLayerCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string

	cmd := &cobra.Command{
		Use:   "register-el",
		Short: "Register as an operator to eigen layer",
		Long:  ``,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}

			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/operator", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}

			if signerConfig.Address == "" {
				logger.Fatal("No signer provided, run with read-only mode")
			}

			privateKey, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
			if err != nil {
				logger.Fatal("Invalid keystore for signer", signerConfig)
				return err
			}

			// NOTE: sanest default for operators, might set it configurable in the futures
			operatorDetails := bindingIDelegationManager.IDelegationManagerOperatorDetails{
				EarningsReceiver:         common.HexToAddress(signerConfig.Address), // earns to self
				DelegationApprover:       common.Address{},                          // address(0x0) means accepts all
				StakerOptOutWindowBlocks: 1,                                         // CAN'T BE DECREASE
			}
			register.RegisterOperatorWithEigenLayer(
				privateKey,
				envConfig,
				operatorDetails,
			)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)

	return cmd
}
