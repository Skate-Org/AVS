package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGRPCError(code codes.Code, msg string) error {
	return status.Errorf(code, msg)
}

// HTTP Mapping: 400 Bad Request
func NewInvalidArgError(msg string) error {
	return NewGRPCError(codes.InvalidArgument, msg)
}

// HTTP Mapping: 404 Not Found
func NewNotFoundError(msg string) error {
	return NewGRPCError(codes.NotFound, msg)
}

// HTTP Mapping: 429 Too Many Requests
func NewResourceExhaustedError(msg string) error {
	return NewGRPCError(codes.ResourceExhausted, msg)
}

// HTTP Mapping: 500 Internal Server Error
func NewInternalError(msg string) error {
	return NewGRPCError(codes.Internal, msg)
}
