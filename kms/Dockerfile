# Start with a Go image as your builder stage
FROM golang:1.21.8-alpine3.18 as builder

# Install necessary tools
RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash

# Copy necessary files and directories into the Docker image
COPY kms /app/kms
COPY lib /app/lib
COPY go.mod /app
COPY go.sum /app

# Set the working directory to the application directory
WORKDIR /app/kms

# Pass version control metadata as build arguments with default values as empty strings
ARG GIT_COMMIT
ARG GIT_TIMESTAMP

# Use build caches to speed up builds
# Build the Go application with ldflags to inject version control metadata
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags "-X 'main.Commit=${GIT_COMMIT}' -X 'main.Timestamp=${GIT_TIMESTAMP}'" -o ./bin/kms .

# Use Alpine for the final image
FROM alpine:3.18

# Copy the built binary from the builder stage to the final image
COPY --from=builder /app/kms/bin/kms /usr/local/bin

# Set the entrypoint to application
ENTRYPOINT ["kms"]

