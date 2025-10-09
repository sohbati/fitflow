#!/bin/bash
# admin-gateway: manage APISIX admin gateway for FitFlow
# Usage: ./admin-gateway.sh [start|stop|restart|status]

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GATEWAY_DIR="$SCRIPT_DIR"
ETCD_PORT=23791
NODE_PORT=9080
ADMIN_PORT=9180
ETCD_DATA="$GATEWAY_DIR/etcd-data"

APISIX_REPO="https://github.com/apache/apisix.git"

# Ensure folder exists
mkdir -p "$GATEWAY_DIR"

# Check if APISIX is installed
if [ ! -d "$GATEWAY_DIR/apisix" ]; then
    echo "[INFO] APISIX not found, downloading..."
    git clone "$APISIX_REPO" "$GATEWAY_DIR/apisix"
    cd "$GATEWAY_DIR/apisix" || exit 1
    make deps
fi

# Check if config exists
CONF_FILE="$GATEWAY_DIR/apisix/conf/config.yaml"
if [ ! -f "$CONF_FILE" ]; then
    echo "[INFO] Config not found, creating default config..."
    cat > "$CONF_FILE" <<EOL
apisix:
  node_listen: $NODE_PORT
  enable_admin: true
  allow_admin:
    - 127.0.0.1
  admin_key:
    - name: admin
      key: admin-key
      role: admin

etcd:
  host:
    - "http://127.0.0.1:$ETCD_PORT"
  prefix: /coach-aid-gateway
  timeout: 30

nginx_config:
  worker_processes: 1
EOL
fi

# Ensure etcd data folder
mkdir -p "$ETCD_DATA"

start() {
    echo "[INFO] Starting CoachAidApp admin-gateway..."
    
    # Start etcd if not running
    if ! pgrep -f "etcd.*$ETCD_DATA" >/dev/null; then
        echo "[INFO] Starting etcd..."
        nohup etcd --data-dir "$ETCD_DATA" \
            --listen-client-urls "http://127.0.0.1:$ETCD_PORT" \
            --advertise-client-urls "http://127.0.0.1:$ETCD_PORT" \
            >/dev/null 2>&1 &
        sleep 2
    fi

    # Start APISIX
    echo "[INFO] Starting APISIX gateway..."
    cd "$GATEWAY_DIR/apisix" || exit 1
    ./bin/apisix start
    echo "[INFO] CoachAidApp admin-gateway started successfully!"
    echo "[INFO] Gateway HTTP port: $NODE_PORT"
    echo "[INFO] Admin API port: $ADMIN_PORT"
    echo "[INFO] Admin key: admin-key"
    echo "[INFO] Gateway URL: http://localhost:$NODE_PORT"
    echo "[INFO] Admin API URL: http://localhost:$ADMIN_PORT"
}

stop() {
    echo "[INFO] Stopping CoachAidApp admin-gateway..."
    cd "$GATEWAY_DIR/apisix" || exit 1
    ./bin/apisix stop

    # Stop etcd
    ETCD_PID=$(pgrep -f "etcd.*$ETCD_DATA")
    if [ -n "$ETCD_PID" ]; then
        kill "$ETCD_PID"
        echo "[INFO] etcd stopped"
    fi
    echo "[INFO] CoachAidApp admin-gateway stopped"
}

status() {
    APISIX_PID=$(pgrep -f "apisix.*nginx")
    ETCD_PID=$(pgrep -f "etcd.*$ETCD_DATA")
    
    echo "=== CoachAidApp Admin Gateway Status ==="
    echo "APISIX PID: ${APISIX_PID:-Not running}"
    echo "ETCD PID: ${ETCD_PID:-Not running}"
    
    if [ -n "$APISIX_PID" ] && [ -n "$ETCD_PID" ]; then
        echo "Status: Running"
        echo "Gateway: http://localhost:$NODE_PORT"
        echo "Admin API: http://localhost:$ADMIN_PORT"
    else
        echo "Status: Stopped"
    fi
}

restart() {
    echo "[INFO] Restarting CoachAidApp admin-gateway..."
    stop
    sleep 2
    start
}

# Check if required tools are installed
check_dependencies() {
    local missing_deps=()
    
    if ! command -v git &> /dev/null; then
        missing_deps+=("git")
    fi
    
    if ! command -v etcd &> /dev/null; then
        missing_deps+=("etcd")
    fi
    
    # Check for macOS-specific dependencies
    if [[ "$OSTYPE" == "darwin"* ]]; then
        if ! command -v brew &> /dev/null; then
            echo "[ERROR] Homebrew is required for macOS installation"
            echo "[INFO] Please install Homebrew: https://brew.sh/"
            exit 1
        fi
        
        # Check for OpenResty dependencies
        if ! brew list luajit-openresty &> /dev/null; then
            missing_deps+=("luajit-openresty")
        fi
        
        if ! brew list lua &> /dev/null; then
            missing_deps+=("lua")
        fi
        
        if ! brew list pcre &> /dev/null; then
            missing_deps+=("pcre")
        fi
        
        if ! brew list pcre2 &> /dev/null; then
            missing_deps+=("pcre2")
        fi
    fi
    
    if [ ${#missing_deps[@]} -ne 0 ]; then
        echo "[ERROR] Missing dependencies: ${missing_deps[*]}"
        echo "[INFO] Please install the missing dependencies:"
        
        if [[ "$OSTYPE" == "darwin"* ]]; then
            echo "  For macOS, run:"
            echo "  brew install ${missing_deps[*]}"
            echo ""
            echo "  Then install OpenResty:"
            echo "  brew install openresty"
        else
            for dep in "${missing_deps[@]}"; do
                case $dep in
                    "git")
                        echo "  - Git: https://git-scm.com/downloads"
                        ;;
                    "etcd")
                        echo "  - etcd: https://etcd.io/docs/v3.5/install/"
                        ;;
                esac
            done
        fi
        exit 1
    fi
}

# Install macOS dependencies
install_macos_deps() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        echo "[INFO] Installing macOS dependencies..."
        
        # Install required packages
        brew install luajit-openresty lua pcre pcre2 go
        
        # Install additional dependencies
        brew install openssl@3
        
        # Add OpenResty to PATH
        export PATH="/usr/local/openresty/nginx/sbin:$PATH"
        
        # Configure LuaRocks for macOS
        local luarocks_config="$HOME/.luarocks/config-5.1.lua"
        if [ ! -f "$luarocks_config" ]; then
            mkdir -p "$(dirname "$luarocks_config")"
            cat > "$luarocks_config" <<EOL
variables = {
   PCRE_DIR = "/usr/local/opt/pcre",
   PCRE_INCDIR = "/usr/local/opt/pcre/include",
   PCRE_LIBDIR = "/usr/local/opt/pcre/lib"
}
EOL
        fi
        
        echo "[INFO] macOS dependencies installed successfully"
    fi
}

# Install dependencies
install() {
    echo "[INFO] Installing CoachAidApp admin-gateway dependencies..."
    
    if [[ "$OSTYPE" == "darwin"* ]]; then
        install_macos_deps
    else
        echo "[INFO] For Linux systems, please install dependencies manually:"
        echo "  - Git: https://git-scm.com/downloads"
        echo "  - etcd: https://etcd.io/docs/v3.5/install/"
        echo "  - OpenResty: https://openresty.org/en/installation.html"
    fi
    
    echo "[INFO] Dependencies installation completed"
}

# Main
case "$1" in
    install)
        install
        ;;
    start) 
        check_dependencies
        start 
        ;;
    stop) 
        stop 
        ;;
    restart) 
        check_dependencies
        restart 
        ;;
    status) 
        status 
        ;;
    *)
        echo "CoachAidApp Admin Gateway Manager"
        echo "Usage: $0 {install|start|stop|restart|status}"
        echo ""
        echo "Commands:"
        echo "  install  - Install required dependencies"
        echo "  start    - Start the admin gateway"
        echo "  stop     - Stop the admin gateway"
        echo "  restart  - Restart the admin gateway"
        echo "  status   - Show gateway status"
        echo ""
        echo "Gateway will be available at:"
        echo "  - Gateway: http://localhost:$NODE_PORT"
        echo "  - Admin API: http://localhost:$ADMIN_PORT"
        echo "  - Admin Key: admin-key"
        echo ""
        echo "First time setup:"
        echo "  ./admin-gateway.sh install"
        echo "  ./admin-gateway.sh start"
        ;;
esac
