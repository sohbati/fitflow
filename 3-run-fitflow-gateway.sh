#!/bin/bash
# Helper to manage the Nginx-based FitFlow gateway

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
GATEWAY_DIR="$ROOT_DIR/fitflow-gateway"
NGINX_CONF="$GATEWAY_DIR/config/nginx.conf"
NGINX_PID_FILE="$GATEWAY_DIR/logs/nginx.pid"
NGINX_ERROR_LOG="$GATEWAY_DIR/logs/error.log"
NGINX_ACCESS_LOG="$GATEWAY_DIR/logs/access.log"

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

find_nginx() {
  # Try common nginx locations
  if command_exists nginx; then
    echo "nginx"
    return 0
  fi
  # Try Homebrew nginx on macOS
  if [ -x "/opt/homebrew/bin/nginx" ]; then
    echo "/opt/homebrew/bin/nginx"
    return 0
  fi
  if [ -x "/usr/local/bin/nginx" ]; then
    echo "/usr/local/bin/nginx"
    return 0
  fi
  return 1
}

require_nginx() {
  local nginx_bin
  nginx_bin=$(find_nginx 2>&1) || true
  if [ -n "$nginx_bin" ] && [ "$nginx_bin" != "" ]; then
    return 0
  fi
  echo "" >&2
  echo "[ERROR] Nginx not found." >&2
  echo "[INFO] Install Nginx:" >&2
  echo "  macOS: brew install nginx" >&2
  echo "  Ubuntu/Debian: sudo apt-get install nginx" >&2
  echo "  CentOS/RHEL: sudo yum install nginx" >&2
  echo "" >&2
  return 1
}

get_nginx_cmd() {
  local nginx_bin
  nginx_bin=$(find_nginx)
  if [ -z "$nginx_bin" ]; then
    echo "[ERROR] Nginx not found" >&2
    exit 1
  fi
  echo "$nginx_bin"
}

ensure_files() {
  if [ ! -f "$NGINX_CONF" ]; then
    echo "[ERROR] Missing Nginx config at $NGINX_CONF" >&2
    exit 1
  fi
  mkdir -p "$GATEWAY_DIR/logs"
  
  # Create mime.types file if it doesn't exist
  if [ ! -f "$GATEWAY_DIR/config/mime.types" ]; then
    local nginx_bin
    nginx_bin=$(get_nginx_cmd)
    # Try to find system mime.types
    local mime_types=""
    for path in /opt/homebrew/etc/nginx/mime.types /usr/local/etc/nginx/mime.types /etc/nginx/mime.types /usr/share/nginx/mime.types; do
      if [ -f "$path" ]; then
        mime_types="$path"
        break
      fi
    done
    
    if [ -n "$mime_types" ]; then
      cp "$mime_types" "$GATEWAY_DIR/config/mime.types"
    else
      # Create a basic mime.types
      cat > "$GATEWAY_DIR/config/mime.types" <<'MIME'
types {
    text/html                             html htm shtml;
    text/css                              css;
    text/xml                              xml;
    image/gif                             gif;
    image/jpeg                            jpeg jpg;
    application/javascript                js;
    application/json                      json;
    application/xml                       xml;
    text/plain                            txt;
}
MIME
    fi
  fi
}

test_config() {
  local nginx_bin
  nginx_bin=$(get_nginx_cmd)
  echo "[INFO] Testing Nginx configuration..."
  echo "[INFO] Config file: $NGINX_CONF"
  echo "[INFO] Working directory: $GATEWAY_DIR"
  "$nginx_bin" -t -c "$NGINX_CONF" -p "$GATEWAY_DIR" || {
    echo "[ERROR] Nginx configuration test failed." >&2
    exit 1
  }
  echo "[INFO] Configuration test passed"
}

extract_listening_ports() {
  # Extract listen directives from nginx config
  local ports
  ports=$(grep -h "listen" "$NGINX_CONF" "$GATEWAY_DIR/config/conf.d"/*.conf 2>/dev/null | \
    grep -v "^\s*#" | \
    grep -v "^\s*$" | \
    sed -E 's/.*listen[[:space:]]+([0-9]+).*/\1/' | \
    grep -E '^[0-9]+$' | \
    sort -un | \
    tr '\n' ' ' | \
    sed 's/ $//')
  echo "$ports"
}

start_gateway() {
  echo "[INFO] ========================================"
  echo "[INFO] Starting FitFlow Nginx Gateway"
  echo "[INFO] ========================================"
  
  if ! require_nginx; then
    exit 1
  fi
  
  ensure_files
  
  local nginx_bin
  nginx_bin=$(get_nginx_cmd)
  echo "[INFO] Nginx binary: $nginx_bin"
  echo "[INFO] Nginx version: $($nginx_bin -v 2>&1 | head -1)"
  
  test_config
  
  # Extract listening ports
  local ports
  ports=$(extract_listening_ports)
  if [ -n "$ports" ]; then
    echo "[INFO] Listening ports: $ports"
  else
    echo "[WARN] Could not determine listening ports from config"
  fi
  
  # Check if already running
  if [ -f "$NGINX_PID_FILE" ]; then
    local pid
    pid=$(cat "$NGINX_PID_FILE" 2>/dev/null || echo "")
    if [ -n "$pid" ] && kill -0 "$pid" 2>/dev/null; then
      echo "[WARN] Nginx is already running (PID: $pid)" >&2
      echo "[INFO] Use './run-fitflow-gateway.sh restart' to restart" >&2
      return 0
    fi
    rm -f "$NGINX_PID_FILE"
  fi
  
  echo "[INFO] Starting Nginx process..."
  "$nginx_bin" -c "$NGINX_CONF" -p "$GATEWAY_DIR"
  sleep 1
  
  if [ -f "$NGINX_PID_FILE" ]; then
    local pid
    pid=$(cat "$NGINX_PID_FILE")
    if kill -0 "$pid" 2>/dev/null; then
      echo "[INFO] ========================================"
      echo "[INFO] Nginx started successfully!"
      echo "[INFO] ========================================"
      echo "[INFO] Process ID: $pid"
      if [ -n "$ports" ]; then
        for port in $ports; do
          echo "[INFO] Listening on port: $port"
          echo "[INFO] Gateway URL: http://localhost:$port"
        done
      else
        echo "[INFO] Gateway URL: http://localhost:8090"
      fi
      echo "[INFO] Config file: $NGINX_CONF"
      echo "[INFO] Error log: $NGINX_ERROR_LOG"
      echo "[INFO] Access log: $NGINX_ACCESS_LOG"
      echo "[INFO] PID file: $NGINX_PID_FILE"
      echo "[INFO] ========================================"
    else
      echo "[ERROR] Nginx process died immediately after start" >&2
      echo "[ERROR] Check error log: $NGINX_ERROR_LOG" >&2
      if [ -f "$NGINX_ERROR_LOG" ]; then
        echo "[ERROR] Last few lines of error log:" >&2
        tail -10 "$NGINX_ERROR_LOG" >&2
      fi
      exit 1
    fi
  else
    echo "[ERROR] Nginx failed to start (no PID file created)" >&2
    echo "[ERROR] Check error log: $NGINX_ERROR_LOG" >&2
    if [ -f "$NGINX_ERROR_LOG" ]; then
      echo "[ERROR] Last few lines of error log:" >&2
      tail -10 "$NGINX_ERROR_LOG" >&2
    fi
    exit 1
  fi
}

stop_gateway() {
  echo "[INFO] Stopping FitFlow Nginx gateway..."
  
  if [ ! -f "$NGINX_PID_FILE" ]; then
    echo "[INFO] Nginx is not running (no PID file found)" >&2
    return 0
  fi
  
  local pid
  pid=$(cat "$NGINX_PID_FILE" 2>/dev/null || echo "")
  
  if [ -z "$pid" ]; then
    echo "[INFO] Nginx is not running" >&2
    rm -f "$NGINX_PID_FILE"
    return 0
  fi
  
  if ! kill -0 "$pid" 2>/dev/null; then
    echo "[INFO] Nginx process not found (PID: $pid)" >&2
    rm -f "$NGINX_PID_FILE"
    return 0
  fi
  
  local nginx_bin
  nginx_bin=$(get_nginx_cmd)
  "$nginx_bin" -s quit -c "$NGINX_CONF" -p "$GATEWAY_DIR" 2>/dev/null || {
    echo "[WARN] Graceful shutdown failed, forcing stop..." >&2
    kill "$pid" 2>/dev/null || true
  }
  
  # Wait for process to stop
  local count=0
  while kill -0 "$pid" 2>/dev/null && [ $count -lt 10 ]; do
    sleep 0.5
    count=$((count + 1))
  done
  
  if kill -0 "$pid" 2>/dev/null; then
    echo "[WARN] Process still running, killing..." >&2
    kill -9 "$pid" 2>/dev/null || true
  fi
  
  rm -f "$NGINX_PID_FILE"
  echo "[INFO] Nginx stopped"
}

status_gateway() {
  if [ ! -f "$NGINX_PID_FILE" ]; then
    echo "[INFO] Nginx is not running (no PID file)" >&2
    exit 1
  fi
  
  local pid
  pid=$(cat "$NGINX_PID_FILE" 2>/dev/null || echo "")
  
  if [ -z "$pid" ]; then
    echo "[INFO] Nginx is not running (empty PID file)" >&2
    exit 1
  fi
  
  if kill -0 "$pid" 2>/dev/null; then
    local ports
    ports=$(extract_listening_ports)
    echo "[INFO] ========================================"
    echo "[INFO] Nginx Gateway Status: RUNNING"
    echo "[INFO] ========================================"
    echo "[INFO] Process ID: $pid"
    if [ -n "$ports" ]; then
      for port in $ports; do
        echo "[INFO] Listening on port: $port"
        echo "[INFO] Gateway URL: http://localhost:$port"
      done
    else
      echo "[INFO] Gateway URL: http://localhost:8090"
    fi
    echo "[INFO] Config file: $NGINX_CONF"
    echo "[INFO] Error log: $NGINX_ERROR_LOG"
    echo "[INFO] Access log: $NGINX_ACCESS_LOG"
    echo "[INFO] ========================================"
    return 0
  else
    echo "[INFO] Nginx is not running (stale PID file: $pid)" >&2
    rm -f "$NGINX_PID_FILE"
    exit 1
  fi
}

logs_gateway() {
  local log_file="${1:-$NGINX_ERROR_LOG}"
  
  if [ ! -f "$log_file" ]; then
    echo "[WARN] Log file $log_file not found." >&2
    return
  fi
  
  tail -f "$log_file"
}

reload_gateway() {
  echo "[INFO] Reloading Nginx configuration..."
  require_nginx
  ensure_files
  
  if [ ! -f "$NGINX_PID_FILE" ]; then
    echo "[ERROR] Nginx is not running. Start it first with './run-fitflow-gateway.sh start'" >&2
    exit 1
  fi
  
  local pid
  pid=$(cat "$NGINX_PID_FILE" 2>/dev/null || echo "")
  if [ -z "$pid" ] || ! kill -0 "$pid" 2>/dev/null; then
    echo "[ERROR] Nginx is not running (PID: $pid)" >&2
    exit 1
  fi
  
  echo "[INFO] Testing configuration before reload..."
  test_config
  
  local nginx_bin
  nginx_bin=$(get_nginx_cmd)
  local ports
  ports=$(extract_listening_ports)
  
  echo "[INFO] Reloading Nginx (PID: $pid)..."
  "$nginx_bin" -s reload -c "$NGINX_CONF" -p "$GATEWAY_DIR"
  
  echo "[INFO] ========================================"
  echo "[INFO] Nginx configuration reloaded successfully"
  echo "[INFO] ========================================"
  if [ -n "$ports" ]; then
    for port in $ports; do
      echo "[INFO] Listening on port: $port"
    done
  fi
  echo "[INFO] Config file: $NGINX_CONF"
  echo "[INFO] ========================================"
}

case "${1:-start}" in
  start)
    start_gateway
    ;;
  stop)
    stop_gateway
    ;;
  restart)
    stop_gateway || true
    sleep 1
    start_gateway
    ;;
  status)
    status_gateway
    ;;
  logs)
    logs_gateway "${2:-}"
    ;;
  reload)
    reload_gateway
    ;;
  test)
    require_nginx
    ensure_files
    test_config
    echo "[INFO] Configuration test passed"
    ;;
  *)
    cat <<USAGE
FitFlow Nginx Gateway Manager
Usage: $0 {start|stop|restart|status|logs|reload|test}

Commands:
  start    Start Nginx gateway using config/nginx.conf
  stop     Stop the Nginx instance
  restart  Restart Nginx
  status   Show Nginx status (fails if not running)
  logs     Tail Nginx logs (default: error.log, accepts optional log file path)
  reload   Hot reload Nginx after editing config files
  test     Test Nginx configuration without starting
USAGE
    exit 1
    ;;
esac
