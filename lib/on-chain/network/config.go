package network

const (
	// Devnet = single-machine deployment, consist of multiple Docker containers.
	Devnet NetworkID = "devnet"

	Testnet NetworkID = "testnet"

	Mainnet NetworkID = "mainnet"
)

var supported = map[NetworkID]bool{
	Devnet:  true,
	Testnet: true,
	Mainnet: true,
}
