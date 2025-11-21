#!/bin/bash

# Run the FitFlow Business service

# Build the FitFlow Business service
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR/fitflow-business"
go build -o "$SCRIPT_DIR/fitflow-business/fitflow-business-bin" ./cmd/main.go

# Run the FitFlow Business service
cd "$SCRIPT_DIR"
./fitflow-business/fitflow-business-bin


