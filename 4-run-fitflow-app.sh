#!/bin/bash
# Helper to manage the FitFlow Next.js application

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
APP_DIR="$ROOT_DIR/fitflow-app"
PID_FILE="$APP_DIR/.next/server.pid"
LOG_FILE="$APP_DIR/logs/app.log"
DEFAULT_PORT=3000

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

require_node() {
  if ! command_exists node; then
    echo "[ERROR] Node.js not found." >&2
    echo "[INFO] Install Node.js:" >&2
    echo "  macOS: brew install node" >&2
    echo "  Or download from: https://nodejs.org/" >&2
    exit 1
  fi
  
  local node_version
  node_version=$(node --version)
  echo "[INFO] Node.js version: $node_version"
  
  if ! command_exists npm; then
    echo "[ERROR] npm not found." >&2
    exit 1
  fi
  
  local npm_version
  npm_version=$(npm --version)
  echo "[INFO] npm version: $npm_version"
}

check_dependencies() {
  if [ ! -d "$APP_DIR/node_modules" ]; then
    echo "[INFO] Dependencies not installed. Installing..."
    cd "$APP_DIR"
    npm install
    echo "[INFO] Dependencies installed"
  fi
}

check_env_file() {
  local env_file="$APP_DIR/.env.local"
  local env_example="$APP_DIR/env.example"
  
  if [ ! -f "$env_file" ]; then
    if [ -f "$env_example" ]; then
      echo "[WARN] .env.local not found. Creating from env.example..." >&2
      cp "$env_example" "$env_file"
      echo "[WARN] Please edit $env_file with your configuration" >&2
      echo "[WARN] Required variables:" >&2
      echo "  - NEXTAUTH_URL" >&2
      echo "  - NEXTAUTH_SECRET" >&2
      echo "  - GOOGLE_CLIENT_ID" >&2
      echo "  - GOOGLE_CLIENT_SECRET" >&2
      echo "" >&2
    else
      echo "[WARN] No .env.local file found. Some features may not work." >&2
    fi
  fi
}

extract_port() {
  local port="${PORT:-$DEFAULT_PORT}"
  if [ -f "$APP_DIR/.env.local" ]; then
    local env_port
    env_port=$(grep -E "^NEXTAUTH_URL=" "$APP_DIR/.env.local" 2>/dev/null | sed -E 's|.*:([0-9]+).*|\1|' || echo "")
    if [ -n "$env_port" ]; then
      port="$env_port"
    fi
  fi
  echo "$port"
}

is_running() {
  local port
  port=$(extract_port)
  
  # Check if process is running on the port
  if command_exists lsof; then
    lsof -ti:$port >/dev/null 2>&1
  elif command_exists netstat; then
    netstat -an | grep -q ":$port.*LISTEN" 2>/dev/null
  else
    # Fallback: check if .next directory exists and is recent
    [ -d "$APP_DIR/.next" ]
  fi
}

get_pid() {
  local port
  port=$(extract_port)
  
  if command_exists lsof; then
    lsof -ti:$port 2>/dev/null | head -1 || echo ""
  else
    # Try to find node process running next
    pgrep -f "next.*dev\|next.*start" 2>/dev/null | head -1 || echo ""
  fi
}

get_all_pids() {
  local port
  port=$(extract_port)
  local pids=""
  
  # Get all PIDs using the port
  if command_exists lsof; then
    pids=$(lsof -ti:$port 2>/dev/null | tr '\n' ' ' | sed 's/ $//')
  fi
  
  # Also find Next.js processes in the app directory
  local next_pids
  next_pids=$(pgrep -f "next.*$APP_DIR\|node.*$APP_DIR.*next" 2>/dev/null | tr '\n' ' ' | sed 's/ $//')
  
  # Combine and deduplicate
  if [ -n "$pids" ] && [ -n "$next_pids" ]; then
    echo "$pids $next_pids" | tr ' ' '\n' | sort -u | tr '\n' ' ' | sed 's/ $//'
  elif [ -n "$pids" ]; then
    echo "$pids"
  elif [ -n "$next_pids" ]; then
    echo "$next_pids"
  else
    echo ""
  fi
}

start_app() {
  echo "[INFO] ========================================"
  echo "[INFO] Starting FitFlow Next.js Application"
  echo "[INFO] ========================================"
  
  require_node
  check_dependencies
  check_env_file
  
  local port
  port=$(extract_port)
  
  if is_running; then
    local pid
    pid=$(get_pid)
    echo "[WARN] Application is already running on port $port (PID: $pid)" >&2
    echo "[INFO] Use './run-fitflow-app.sh restart' to restart" >&2
    return 0
  fi
  
  cd "$APP_DIR"
  mkdir -p "$(dirname "$LOG_FILE")"
  
  echo "[INFO] Starting Next.js development server..."
  echo "[INFO] Port: $port"
  echo "[INFO] Working directory: $APP_DIR"
  
  # Start in background and capture PID
  npm run dev > "$LOG_FILE" 2>&1 &
  local pid=$!
  
  # Wait a moment for server to start
  sleep 3
  
  if is_running; then
    echo "[INFO] ========================================"
    echo "[INFO] FitFlow App started successfully!"
    echo "[INFO] ========================================"
    echo "[INFO] Process ID: $pid"
    echo "[INFO] Listening on port: $port"
    echo "[INFO] Application URL: http://localhost:$port"
    echo "[INFO] Log file: $LOG_FILE"
    echo "[INFO] Working directory: $APP_DIR"
    echo "[INFO] ========================================"
    echo "[INFO] To view logs: ./run-fitflow-app.sh logs"
    echo "[INFO] To stop: ./run-fitflow-app.sh stop"
    echo "[INFO] ========================================"
  else
    echo "[ERROR] Application failed to start" >&2
    echo "[ERROR] Check log file: $LOG_FILE" >&2
    if [ -f "$LOG_FILE" ]; then
      echo "[ERROR] Last few lines of log:" >&2
      tail -20 "$LOG_FILE" >&2
    fi
    exit 1
  fi
}

stop_app() {
  echo "[INFO] Stopping FitFlow Next.js Application..."
  
  local port
  port=$(extract_port)
  
  # Get all related PIDs
  local all_pids
  all_pids=$(get_all_pids)
  
  if [ -z "$all_pids" ] && ! is_running; then
    echo "[INFO] Application is not running on port $port" >&2
    return 0
  fi
  
  if [ -z "$all_pids" ]; then
    echo "[WARN] Could not find process IDs, but port $port appears to be in use" >&2
    echo "[INFO] Attempting to find and kill processes on port $port..." >&2
    
    if command_exists lsof; then
      all_pids=$(lsof -ti:$port 2>/dev/null | tr '\n' ' ' | sed 's/ $//')
    fi
  fi
  
  if [ -z "$all_pids" ]; then
    echo "[INFO] No processes found to stop" >&2
    return 0
  fi
  
  echo "[INFO] Found processes: $all_pids"
  echo "[INFO] Stopping processes..."
  
  # Kill all processes
  for pid in $all_pids; do
    if kill -0 "$pid" 2>/dev/null; then
      echo "[INFO] Stopping process $pid..."
      kill "$pid" 2>/dev/null || true
    fi
  done
  
  # Wait for processes to stop
  local count=0
  local still_running=""
  while [ $count -lt 10 ]; do
    still_running=""
    for pid in $all_pids; do
      if kill -0 "$pid" 2>/dev/null; then
        still_running="$still_running $pid"
      fi
    done
    
    if [ -z "$still_running" ]; then
      break
    fi
    
    sleep 0.5
    count=$((count + 1))
  done
  
  # Force kill any remaining processes
  if [ -n "$still_running" ]; then
    echo "[WARN] Some processes still running, forcing stop..." >&2
    for pid in $still_running; do
      if kill -0 "$pid" 2>/dev/null; then
        echo "[INFO] Force killing process $pid..."
        kill -9 "$pid" 2>/dev/null || true
      fi
    done
    sleep 1
  fi
  
  # Verify port is free
  if is_running; then
    echo "[WARN] Port $port may still be in use. Checking..." >&2
    if command_exists lsof; then
      local remaining
      remaining=$(lsof -ti:$port 2>/dev/null | tr '\n' ' ' | sed 's/ $//')
      if [ -n "$remaining" ]; then
        echo "[WARN] Port $port still in use by: $remaining" >&2
        echo "[INFO] You may need to manually kill these processes" >&2
      fi
    fi
  else
    echo "[INFO] Application stopped successfully"
  fi
}

status_app() {
  local port
  port=$(extract_port)
  
  if ! is_running; then
    echo "[INFO] Application is not running on port $port" >&2
    exit 1
  fi
  
  local pid
  pid=$(get_pid)
  
  echo "[INFO] ========================================"
  echo "[INFO] FitFlow App Status: RUNNING"
  echo "[INFO] ========================================"
  echo "[INFO] Process ID: $pid"
  echo "[INFO] Listening on port: $port"
  echo "[INFO] Application URL: http://localhost:$port"
  echo "[INFO] Working directory: $APP_DIR"
  echo "[INFO] Log file: $LOG_FILE"
  echo "[INFO] ========================================"
}

logs_app() {
  local log_file="${1:-$LOG_FILE}"
  
  if [ ! -f "$log_file" ]; then
    echo "[WARN] Log file $log_file not found." >&2
    echo "[INFO] Application may not be running or hasn't generated logs yet." >&2
    return
  fi
  
  echo "[INFO] Tailing log file: $log_file"
  echo "[INFO] Press Ctrl+C to stop"
  echo ""
  tail -f "$log_file"
}

restart_app() {
  stop_app || true
  sleep 2
  start_app
}

build_app() {
  echo "[INFO] Building FitFlow Next.js Application..."
  require_node
  check_dependencies
  
  cd "$APP_DIR"
  npm run build
  echo "[INFO] Build completed"
}

case "${1:-start}" in
  start)
    start_app
    ;;
  stop)
    stop_app
    ;;
  restart)
    restart_app
    ;;
  status)
    status_app
    ;;
  logs)
    logs_app "${2:-}"
    ;;
  build)
    build_app
    ;;
  install)
    require_node
    check_dependencies
    ;;
  *)
    cat <<USAGE
FitFlow Next.js Application Manager
Usage: $0 {start|stop|restart|status|logs|build|install}

Commands:
  start    Start the Next.js development server
  stop     Stop the running application
  restart  Restart the application
  status   Show application status
  logs     Tail application logs (accepts optional log file path)
  build    Build the application for production
  install  Install npm dependencies

Environment:
  The app runs on port 3000 by default (or PORT env var)
  Configuration: $APP_DIR/.env.local
  Logs: $LOG_FILE
USAGE
    exit 1
    ;;
esac

