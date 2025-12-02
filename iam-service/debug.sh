#!/bin/bash
# Debug script for IAM Service using Delve

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# Check if Delve is installed
if ! command -v dlv &> /dev/null; then
    echo "Delve not found. Installing..."
    go install github.com/go-delve/delve/cmd/dlv@latest
fi

# Load environment variables from .env if it exists
if [ -f ".env" ]; then
    export $(cat .env | grep -v '^#' | xargs)
    echo "[INFO] Loaded environment variables from .env"
fi

# Set default values if not set
export JWT_SECRET=${JWT_SECRET:-mysecret}
export DATABASE_URL=${DATABASE_URL:-postgres://fitflow_iam_user:password@localhost:5432/fitflow_iam_db?sslmode=disable}
export DATABASE_TYPE=${DATABASE_TYPE:-postgres}
export TOKEN_EXP_MINUTES=${TOKEN_EXP_MINUTES:-15}
export PORT=${PORT:-8091}
export GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID:-}
export GOOGLE_CLIENT_SECRET=${GOOGLE_CLIENT_SECRET:-}
export GOOGLE_REDIRECT_URL=${GOOGLE_REDIRECT_URL:-http://localhost:3000/auth/google/callback}

echo "[INFO] Starting IAM Service in debug mode..."
echo "[INFO] Port: $PORT"
echo "[INFO] Google OAuth configured: $([ -n "$GOOGLE_CLIENT_ID" ] && echo "Yes" || echo "No")"
echo "[INFO] Breakpoints will work. Connect your debugger or use: dlv connect localhost:2345"
echo ""

# Run with Delve
dlv debug ./cmd/main.go --listen=:2345 --headless=false --api-version=2 --check-go-version=false

