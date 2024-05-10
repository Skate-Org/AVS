package network

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/Skate-Org/AVS/lib/logging"

	"github.com/ethereum/go-ethereum/common"
)

type ChainID = uint32

// EVMChain defines the configuration of an execution chain that supports
type EVMChain struct {
	ID            ChainID        // Chain ID asa per https://chainlist.org
	Name          string         // Chain name as per https://chainlist.org
	RPCURL        string         // RPC URL of the chain
	AuthRPCURL    string         // RPC URL of the chain with JWT authentication enabled
	PortalAddress common.Address // Address of the omni portal contract on the chain
	DeployHeight  uint64         // Height that the portal contracts were deployed
	IsEthereum    bool           // Whether this is the ethereum layer1 chain
	BlockDuration time.Duration  // Block period of the chain
}

// Load loads the network configuration from the given path.
func Load(path string) (Network, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return Network{}, errors.Wrap(err, "read network config file")
	}

	var net Network
	if err := json.Unmarshal(bz, &net); err != nil {
		return Network{}, errors.Wrap(err, "unmarshal network config file")
	}

	return net, nil
}

// Save saves the network configuration to the given path.
func Save(ctx context.Context, network Network, path string) error {
	logger := logging.NewLoggerWithConsoleWriter()
	for _, chain := range network.Chains {
		if chain.PortalAddress == (common.Address{}) {
			logger.Warn("Netconf network.json portal address empty", "chain", chain.Name, "path", path)
		}
	}

	bz, err := json.MarshalIndent(network, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal network config file")
	}

	if err := os.WriteFile(path, bz, 0o600); err != nil {
		return errors.Wrap(err, "write network config file")
	}

	return nil
}

type chainJSON struct {
	ID            uint32 `json:"id"`
	Name          string `json:"name"`
	RPCURL        string `json:"rpcurl"`
	AuthRPCURL    string `json:"auth_rpcurl,omitempty"`
	PortalAddress string `json:"portal_address"`
	DeployHeight  uint64 `json:"deploy_height"`
	IsEthereum    bool   `json:"is_ethereum,omitempty"`
	BlockPeriod   string `json:"block_period"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *EVMChain) UnmarshalJSON(bz []byte) error {
	var cj chainJSON
	if err := json.Unmarshal(bz, &cj); err != nil {
		return errors.Wrap(err, "unmarshal chain")
	}

	blockPeriod, err := time.ParseDuration(cj.BlockPeriod)
	if err != nil {
		return errors.Wrap(err, "parse block period")
	}

	var portalAddr common.Address
	if cj.PortalAddress != "" {
		portalAddr = common.HexToAddress(cj.PortalAddress)
	}

	*c = EVMChain{
		ID:            cj.ID,
		Name:          cj.Name,
		RPCURL:        cj.RPCURL,
		AuthRPCURL:    cj.AuthRPCURL,
		PortalAddress: portalAddr,
		DeployHeight:  cj.DeployHeight,
		IsEthereum:    cj.IsEthereum,
		BlockDuration: blockPeriod,
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (c EVMChain) MarshalJSON() ([]byte, error) {
	portalAddr := c.PortalAddress.Hex()
	if c.PortalAddress == (common.Address{}) {
		portalAddr = ""
	}

	cj := chainJSON{
		ID:            c.ID,
		Name:          c.Name,
		RPCURL:        c.RPCURL,
		AuthRPCURL:    c.AuthRPCURL,
		PortalAddress: portalAddr,
		DeployHeight:  c.DeployHeight,
		IsEthereum:    c.IsEthereum,
		BlockPeriod:   c.BlockDuration.String(),
	}

	bz, err := json.Marshal(cj)
	if err != nil {
		return nil, errors.Wrap(err, "marshal chain")
	}

	return bz, nil
}
