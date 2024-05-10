package api

import (
	pb "skatechain.org/api/pb/relayer"
)

func Uint32ToChainType(num uint32) pb.ChainType {
	switch num {
	case 0:
		return pb.ChainType_EVM
	default:
		panic("Invalid conversion")
	}
}
