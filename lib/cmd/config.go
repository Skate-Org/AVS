package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	// Level
	Environment string `mapstructure:"environment"`

	// Skatechain
	SkateChainId uint64 `mapstructure:"skate_chain_id"`
	SkateWSSRPC  string `mapstructure:"skate_wss_rpc"`
	SkateHttpRPC string `mapstructure:"skate_http_rpc"`
	SkateApp     string `mapstructure:"skate_app"`

	// Eigenlayer
	MainChainId uint64 `mapstructure:"main_chain_id"`
	SkateAVS          string `mapstructure:"skate_avs"`
	WsETHStrategy     string `mapstructure:"wsETH_strategy"`
	DelegationManager string `mapstructure:"delegation_manager"`
	StrategyManager   string `mapstructure:"strategy_manager"`
	AVSDirectory      string `mapstructure:"avs_directory"`
	// AVS network
	HttpRPC string `mapstructure:"http_rpc"`
	WsRPC   string `mapstructure:"wss_rpc"`
	// Strategies
	Strategy_stETH string `mapstructure:"strategy_stETH"`
	Token_stETH    string `mapstructure:"token_stETH"`

	// AVS node configs
	EigenMetricsIPPort string `mapstructure:"eigen_metrics_ip_port_address"`
	EnableMetrics      bool   `mapstructure:"enable_metrics"`
	NodeAPIIPPort      string `mapstructure:"node_api_ip_port_address"`
	EnableNodeAPI      bool   `mapstructure:"enable_node_api"`
}

type SignerConfig struct {
	Address    string `mapstructure:"address"`
	Passphrase string `mapstructure:"passphrase"`
}

func ReadConfig[T any](subDir string, filename string) (*T, error) {
	// Set the name of the config file (without extension)
	viper.SetConfigName(filename)
	// Set the path to look for the config file
	viper.AddConfigPath(fmt.Sprintf("configs%s", subDir))
	// Set the config file type
	viper.SetConfigType("yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	// Initialize a Config struct to hold the config values
	var config T

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}
	return &config, nil
}
