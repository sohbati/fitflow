#!/bin/bash

# Helper runner for the Admin Gateway from repository root.
# Usage: ./run-admin-gateway.sh [install|start|stop|restart|status]

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
GATEWAY_SCRIPT="$SCRIPT_DIR/admin-gateway/admin-gateway-macos.sh"

if [ ! -f "$GATEWAY_SCRIPT" ]; then
  echo "Error: expected gateway script at $GATEWAY_SCRIPT" >&2
  echo "Please ensure the admin-gateway project is checked out." >&2
  exit 1
fi

ACTION="${1:-start}"

(
  cd "$SCRIPT_DIR/admin-gateway"
  bash "./$(basename "$GATEWAY_SCRIPT")" "$ACTION"
)

