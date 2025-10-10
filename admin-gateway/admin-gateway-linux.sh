#!/bin/bash
# FitFlow APISIX Linux Deployment Script
# Usage: ./deploy-linux.sh [ubuntu|alpine] [production|development]

set -e

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG_DIR="$SCRIPT_DIR/config"
APISIX_VERSION="3.8.0"
ETCD_VERSION="3.5.7"

# Default values
OS_TYPE="${1:-ubuntu}"
ENV_TYPE="${2:-production}"
APISIX_DIR="/opt/apisix"
ETCD_DIR="/opt/etcd"
CONFIG_FILE="$CONFIG_DIR/config.yaml"
ENV_FILE="$CONFIG_DIR/env.$ENV_TYPE"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root
check_root() {
    if [[ $EUID -ne 0 ]]; then
        log_error "This script must be run as root"
        exit 1
    fi
}

# Install dependencies for Ubuntu
install_ubuntu_deps() {
    log_info "Installing dependencies for Ubuntu..."
    
    apt-get update
    apt-get install -y \
        wget \
        curl \
        git \
        build-essential \
        libpcre3-dev \
        zlib1g-dev \
        libssl-dev \
        unzip \
        supervisor \
        nginx-common
    
    # Install OpenResty
    wget -O - https://openresty.org/package/pubkey.gpg | apt-key add -
    echo "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main" > /etc/apt/sources.list.d/openresty.list
    apt-get update
    apt-get install -y openresty
    
    # Install etcd
    wget https://github.com/etcd-io/etcd/releases/download/v${ETCD_VERSION}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz
    tar -xzf etcd-v${ETCD_VERSION}-linux-amd64.tar.gz
    mv etcd-v${ETCD_VERSION}-linux-amd64/etcd* /usr/local/bin/
    rm -rf etcd-v${ETCD_VERSION}-linux-amd64*
    
    # Install APISIX
    wget https://github.com/apache/apisix/releases/download/${APISIX_VERSION}/apisix-${APISIX_VERSION}-src.tar.gz
    tar -xzf apisix-${APISIX_VERSION}-src.tar.gz
    mv apisix-${APISIX_VERSION} $APISIX_DIR
    cd $APISIX_DIR
    make deps
    
    log_info "Ubuntu dependencies installed successfully"
}

# Install dependencies for Alpine
install_alpine_deps() {
    log_info "Installing dependencies for Alpine..."
    
    apk update
    apk add --no-cache \
        wget \
        curl \
        git \
        build-base \
        pcre-dev \
        zlib-dev \
        openssl-dev \
        unzip \
        supervisor \
        nginx
    
    # Install OpenResty
    wget https://openresty.org/download/openresty-1.21.4.3.tar.gz
    tar -xzf openresty-1.21.4.3.tar.gz
    cd openresty-1.21.4.3
    ./configure --with-pcre-jit --with-ipv6 --with-http_realip_module --with-http_ssl_module --with-http_stub_status_module --with-http_gzip_static_module --with-luajit
    make -j$(nproc)
    make install
    cd ..
    rm -rf openresty-1.21.4.3*
    
    # Install etcd
    wget https://github.com/etcd-io/etcd/releases/download/v${ETCD_VERSION}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz
    tar -xzf etcd-v${ETCD_VERSION}-linux-amd64.tar.gz
    mv etcd-v${ETCD_VERSION}-linux-amd64/etcd* /usr/local/bin/
    rm -rf etcd-v${ETCD_VERSION}-linux-amd64*
    
    # Install APISIX
    wget https://github.com/apache/apisix/releases/download/${APISIX_VERSION}/apisix-${APISIX_VERSION}-src.tar.gz
    tar -xzf apisix-${APISIX_VERSION}-src.tar.gz
    mv apisix-${APISIX_VERSION} $APISIX_DIR
    cd $APISIX_DIR
    make deps
    
    log_info "Alpine dependencies installed successfully"
}

# Create systemd services
create_systemd_services() {
    log_info "Creating systemd services..."
    
    # ETCD service
    cat > /etc/systemd/system/etcd.service <<EOF
[Unit]
Description=etcd
After=network.target

[Service]
Type=notify
ExecStart=/usr/local/bin/etcd --data-dir=/var/lib/etcd --listen-client-urls=http://127.0.0.1:2379 --advertise-client-urls=http://127.0.0.1:2379
Restart=always
RestartSec=5
User=etcd
Group=etcd

[Install]
WantedBy=multi-user.target
EOF

    # APISIX service
    cat > /etc/systemd/system/apisix.service <<EOF
[Unit]
Description=APISIX
After=etcd.service
Requires=etcd.service

[Service]
Type=forking
ExecStart=$APISIX_DIR/bin/apisix start
ExecStop=$APISIX_DIR/bin/apisix stop
ExecReload=$APISIX_DIR/bin/apisix reload
Restart=always
RestartSec=5
User=apisix
Group=apisix

[Install]
WantedBy=multi-user.target
EOF

    # Create users
    useradd -r -s /bin/false apisix
    useradd -r -s /bin/false etcd
    
    # Create directories
    mkdir -p /var/lib/etcd
    mkdir -p /var/log/apisix
    mkdir -p /var/run/apisix
    chown -R apisix:apisix /var/log/apisix /var/run/apisix
    chown -R etcd:etcd /var/lib/etcd
    
    systemctl daemon-reload
    systemctl enable etcd apisix
    
    log_info "Systemd services created successfully"
}

# Copy configuration files
copy_config() {
    log_info "Copying configuration files..."
    
    if [ ! -f "$CONFIG_FILE" ]; then
        log_error "Configuration file not found: $CONFIG_FILE"
        exit 1
    fi
    
    cp "$CONFIG_FILE" "$APISIX_DIR/conf/config.yaml"
    chown apisix:apisix "$APISIX_DIR/conf/config.yaml"
    
    # Copy environment file if it exists
    if [ -f "$ENV_FILE" ]; then
        cp "$ENV_FILE" "/etc/apisix.env"
        chown apisix:apisix "/etc/apisix.env"
        log_info "Environment file copied: $ENV_FILE"
    else
        log_warn "Environment file not found: $ENV_FILE"
        log_info "Using default environment settings"
    fi
    
    log_info "Configuration files copied successfully"
}

# Create environment file
create_env_file() {
    log_info "Creating environment file..."
    
    cat > /etc/apisix.env <<EOF
# CoachAidApp APISIX Environment Configuration
APISIX_ADMIN_KEY=coach-aid-admin-key-$(date +%Y%m%d)
ETCD_HOST=http://127.0.0.1:2379
APISIX_NODE_PORT=9080
APISIX_ALLOW_ADMIN=127.0.0.1
NGINX_WORKER_PROCESSES=auto
EOF
    
    chown apisix:apisix /etc/apisix.env
    
    log_info "Environment file created successfully"
}

# Start services
start_services() {
    log_info "Starting services..."
    
    systemctl start etcd
    sleep 5
    systemctl start apisix
    
    log_info "Services started successfully"
}

# Show status
show_status() {
    log_info "Service Status:"
    echo "ETCD: $(systemctl is-active etcd)"
    echo "APISIX: $(systemctl is-active apisix)"
    echo ""
    echo "Gateway URL: http://localhost:9080"
    echo "Admin API URL: http://localhost:9180"
    echo "Admin Key: $(grep APISIX_ADMIN_KEY /etc/apisix.env | cut -d'=' -f2)"
}

# Main deployment function
deploy() {
    log_info "Starting CoachAidApp APISIX deployment..."
    log_info "OS: $OS_TYPE, Environment: $ENV_TYPE"
    
    check_root
    
    case $OS_TYPE in
        ubuntu)
            install_ubuntu_deps
            ;;
        alpine)
            install_alpine_deps
            ;;
        *)
            log_error "Unsupported OS type: $OS_TYPE"
            log_error "Supported types: ubuntu, alpine"
            exit 1
            ;;
    esac
    
    create_systemd_services
    copy_config
    create_env_file
    start_services
    show_status
    
    log_info "Deployment completed successfully!"
}

# Show usage
usage() {
    echo "CoachAidApp APISIX Linux Deployment Script"
    echo ""
    echo "Usage: $0 [OS_TYPE] [ENVIRONMENT]"
    echo ""
    echo "OS_TYPE:"
    echo "  ubuntu  - Deploy on Ubuntu/Debian (default)"
    echo "  alpine  - Deploy on Alpine Linux"
    echo ""
    echo "ENVIRONMENT:"
    echo "  production  - Production configuration (default)"
    echo "  development - Development configuration"
    echo ""
    echo "Examples:"
    echo "  $0 ubuntu production"
    echo "  $0 alpine development"
    echo ""
    echo "Prerequisites:"
    echo "  - Run as root"
    echo "  - Internet connection"
    echo "  - At least 2GB RAM"
    echo "  - At least 10GB disk space"
}

# Main
case "${1:-help}" in
    ubuntu|alpine)
        deploy
        ;;
    help|--help|-h)
        usage
        ;;
    *)
        log_error "Invalid arguments"
        usage
        exit 1
        ;;
esac
