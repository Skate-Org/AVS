package network

import (
	"fmt"

	"github.com/pkg/errors"
)

// TODO: extends this package for non-EVM chain

// Network defines a deployment of Skate avs.
type Network struct {
	ID     NetworkID  `json:"name"`   // ID of the network. e.g. "simnet", "testnet", "staging", "mainnet"
	Chains []EVMChain `json:"chains"` // Chains that are part of the network
}

// Validate returns an error if the configuration is invalid.
func (n Network) Validate() error {
	if err := n.ID.Validate(); err != nil {
		return err
	}
	// NOTE: add more validations in futures for non-EVM
	return nil
}

// ChainIDs returns the all chain IDs in the network.
func (n Network) ChainIDs() []uint32 {
	resp := make([]uint32, 0, len(n.Chains))
	for _, chain := range n.Chains {
		resp = append(resp, chain.ID)
	}

	return resp
}

// ChainNamesByIDs returns the all chain IDs and names in the network.
func (n Network) ChainNamesByIDs() map[uint32]string {
	resp := make(map[uint32]string)
	for _, chain := range n.Chains {
		resp[chain.ID] = chain.Name
	}

	return resp
}

// Chain returns the chain config for the given ID or false if it does not exist.
func (n Network) GetChain(id uint32) (EVMChain, bool) {
	for _, chain := range n.Chains {
		if chain.ID == id {
			return chain, true
		}
	}

	return EVMChain{}, false
}

// NetworkID is a network identifier.
type NetworkID string

// IsProtected returns true if the network is long-lived, therefore protected.
func (i NetworkID) IsPersistent() bool {
	nets := []NetworkID{Devnet}

	for _, net := range nets {
		if i == net {
			return false
		}
	}

	return true
}

func (id NetworkID) Validate() error {
	if !supported[id] {
		return errors.New(fmt.Sprintf("unsupported network: %s", id))
	}

	return nil
}

func (i NetworkID) String() string {
	return string(i)
}
