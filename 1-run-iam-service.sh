#!/bin/bash

# Run the IAM service

# Build the IAM service
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
IAM_SERVICE_DIR="$SCRIPT_DIR/iam-service"
cd "$IAM_SERVICE_DIR"
go build -o "$IAM_SERVICE_DIR/iam-service-bin" ./cmd/main.go

# Run the IAM service from the iam-service directory so .env file is found
cd "$IAM_SERVICE_DIR"
./iam-service-bin


