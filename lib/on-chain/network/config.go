package network

import (
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
)

func IsSupported(chainType uint32) bool {
	switch chainType {
	case uint32(pb.ChainType_EVM), uint32(pb.ChainType_SOLANA):
		return true
	default:
		return false
	}
}
