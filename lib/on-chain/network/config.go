package network

import (
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
)

func IsSupported(chainType uint32, chainId uint32) bool {
	switch chainType {
	case uint32(pb.ChainType_EVM):
		switch chainId {
		case 421614: // Arbitrum Sepolia
			return true
		default:
			return false
		}
	case uint32(pb.ChainType_SOLANA):
		switch chainId {
		case 0: // do not accept other
			return true
		default:
			return false
		}
	default:
		return false
	}
}
