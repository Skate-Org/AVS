package skateapp

import (
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
)

type SignedTask struct {
	ChainType uint32 `dynamodbav:"ChainType"`
	ChainId   uint32 `dynamodbav:"ChainId"`
	TaskId    uint32 `dynamodbav:"TaskId"`
	Message   string `dynamodbav:"Message"`
	Initiator string `dynamodbav:"Initiator"`
	Hash      []byte `dynamodbav:"Hash"`
	Operator  string `dynamodbav:"Operator"`
	Signature []byte `dynamodbav:"Signature"`
}

type CompletedTask struct {
	ChainType uint32 `dynamodbav:"ChainType"`
	ChainId   uint32 `dynamodbav:"ChainId"`
	TaskId    uint32 `dynamodbav:"TaskId"`
	Message   string `dynamodbav:"Message"`
	Initiator string `dynamodbav:"Initiator"`
}

type SignatureTuple struct {
	Operator  string
	Signature []byte
}

type Message struct {
	TaskId    uint32
	Initiator string
	Message   string
	ChainType pb.ChainType
	ChainId   uint32
}
