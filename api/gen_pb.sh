#!/bin/bash

# Directory containing proto files
PROTO_DIR="./proto"
# Output directory for Go files
GO_OUT_DIR="./pb"

rm -rf $GO_OUT_DIR
mkdir -p $GO_OUT_DIR

# Find all .proto files in the PROTO_DIR and iterate over them
find ${PROTO_DIR} -name "*.proto" | while read -r proto_file; do
    # Compute the subdirectory path by removing the PROTO_DIR part from the proto file path
    SUB_DIR=$(dirname "${proto_file/${PROTO_DIR}/}")
    echo "generating pb for " $SUB_DIR
    # Ensure the output directory exists
    mkdir -p "${GO_OUT_DIR}/${SUB_DIR}"
    # Run protoc with the appropriate paths
    protoc -I "${PROTO_DIR}" \
           --go_out="${GO_OUT_DIR}" --go_opt=paths=source_relative \
           --go-grpc_out="${GO_OUT_DIR}" --go-grpc_opt=paths=source_relative \
           "${proto_file}"
    echo "successfully generated pb for" $SUB_DIR
done

echo "Proto files have been compiled into Go files in ${GO_OUT_DIR}"
