#!/bin/bash

# Run the IAM service

# Build the IAM service
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR/iam-service"
go build -o "$SCRIPT_DIR/iam-service/iam-service-bin" ./cmd/main.go

# Run the IAM service
cd "$SCRIPT_DIR"
./iam-service/iam-service-bin


