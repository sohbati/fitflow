#!/bin/bash
# CoachAidApp APISIX Configuration Manager
# Usage: ./config-manager.sh [command] [environment]

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG_DIR="$SCRIPT_DIR/config"
APISIX_DIR="/opt/apisix"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Show available configurations
list_configs() {
    log_info "Available configurations:"
    echo ""
    
    # Show main config
    if [ -f "$CONFIG_DIR/config.yaml" ]; then
        echo "✓ Main Configuration"
        echo "  File: $CONFIG_DIR/config.yaml"
        echo "  Size: $(du -h "$CONFIG_DIR/config.yaml" | cut -f1)"
        echo ""
    else
        echo "✗ Main Configuration (missing)"
        echo ""
    fi
    
    # Show environment files
    echo "Environment Files:"
    for env_file in "$CONFIG_DIR"/env.*; do
        if [ -f "$env_file" ]; then
            env_name=$(basename "$env_file" | sed 's/env\.//')
            echo "✓ $env_name"
            echo "  File: $env_file"
            echo "  Size: $(du -h "$env_file" | cut -f1)"
            echo ""
        fi
    done
}

# Apply configuration
apply_config() {
    local env="$1"
    local config_file="$CONFIG_DIR/config.yaml"
    local env_file="$CONFIG_DIR/env.$env"
    
    if [ ! -f "$config_file" ]; then
        log_error "Configuration file not found: $config_file"
        exit 1
    fi
    
    log_info "Applying configuration with $env environment..."
    
    # Backup current config
    if [ -f "$APISIX_DIR/conf/config.yaml" ]; then
        cp "$APISIX_DIR/conf/config.yaml" "$APISIX_DIR/conf/config.yaml.backup.$(date +%Y%m%d_%H%M%S)"
        log_info "Backed up current configuration"
    fi
    
    # Copy main config
    cp "$config_file" "$APISIX_DIR/conf/config.yaml"
    chown apisix:apisix "$APISIX_DIR/conf/config.yaml" 2>/dev/null || true
    
    # Copy environment file if it exists
    if [ -f "$env_file" ]; then
        cp "$env_file" "/etc/apisix.env"
        chown apisix:apisix "/etc/apisix.env" 2>/dev/null || true
        log_info "Environment file applied: $env_file"
    else
        log_warn "Environment file not found: $env_file"
        log_info "Using default environment settings"
    fi
    
    log_info "Configuration applied successfully"
    
    # Test configuration
    if command -v apisix >/dev/null 2>&1; then
        log_info "Testing configuration..."
        if apisix test; then
            log_info "Configuration test passed"
        else
            log_error "Configuration test failed"
            exit 1
        fi
    fi
}

# Validate configuration
validate_config() {
    local env="$1"
    local config_file="$CONFIG_DIR/config.yaml"
    local env_file="$CONFIG_DIR/env.$env"
    
    if [ ! -f "$config_file" ]; then
        log_error "Configuration file not found: $config_file"
        exit 1
    fi
    
    log_info "Validating configuration with $env environment..."
    
    # Basic YAML syntax check
    if command -v yq >/dev/null 2>&1; then
        if yq eval '.' "$config_file" >/dev/null 2>&1; then
            log_info "YAML syntax is valid"
        else
            log_error "YAML syntax error in $config_file"
            exit 1
        fi
    else
        log_warn "yq not installed, skipping YAML validation"
    fi
    
    # Check required fields
    local required_fields=("apisix.node_listen" "apisix.enable_admin" "deployment.role" "deployment.etcd.host")
    
    for field in "${required_fields[@]}"; do
        if grep -q "$field" "$config_file"; then
            log_info "✓ Found required field: $field"
        else
            log_warn "⚠ Missing field: $field"
        fi
    done
    
    # Check environment file if it exists
    if [ -f "$env_file" ]; then
        log_info "✓ Environment file found: $env_file"
    else
        log_warn "⚠ Environment file not found: $env_file"
    fi
    
    log_info "Configuration validation completed"
}

# Generate environment file
generate_env() {
    local env="$1"
    local env_file="/etc/apisix.env"
    
    log_info "Generating environment file for $env..."
    
    case $env in
        production)
            cat > "$env_file" <<EOF
# CoachAidApp APISIX Production Environment
APISIX_ADMIN_KEY=coach-aid-prod-key-$(date +%Y%m%d)
ETCD_HOST=http://127.0.0.1:2379
APISIX_NODE_PORT=9080
APISIX_ALLOW_ADMIN=127.0.0.1
NGINX_WORKER_PROCESSES=auto
EOF
            ;;
        development)
            cat > "$env_file" <<EOF
# CoachAidApp APISIX Development Environment
APISIX_ADMIN_KEY=coach-aid-dev-key
ETCD_HOST=http://127.0.0.1:23791
APISIX_NODE_PORT=9080
APISIX_ALLOW_ADMIN=0.0.0.0/0
NGINX_WORKER_PROCESSES=1
EOF
            ;;
        *)
            log_error "Unknown environment: $env"
            exit 1
            ;;
    esac
    
    chown apisix:apisix "$env_file" 2>/dev/null || true
    log_info "Environment file generated: $env_file"
}

# Show configuration diff
show_diff() {
    local env="$1"
    local config_file="$CONFIG_DIR/$env/config.yaml"
    
    if [ ! -f "$config_file" ]; then
        log_error "Configuration file not found: $config_file"
        exit 1
    fi
    
    if [ ! -f "$APISIX_DIR/conf/config.yaml" ]; then
        log_info "No current configuration found"
        return
    fi
    
    log_info "Showing diff between $env and current configuration:"
    echo ""
    diff -u "$APISIX_DIR/conf/config.yaml" "$config_file" || true
}

# Restart services
restart_services() {
    log_info "Restarting APISIX services..."
    
    if systemctl is-active --quiet apisix; then
        systemctl restart apisix
        log_info "APISIX restarted"
    else
        log_warn "APISIX service not running"
    fi
    
    if systemctl is-active --quiet etcd; then
        systemctl restart etcd
        log_info "ETCD restarted"
    else
        log_warn "ETCD service not running"
    fi
}

# Show usage
usage() {
    echo "CoachAidApp APISIX Configuration Manager"
    echo ""
    echo "Usage: $0 [command] [environment]"
    echo ""
    echo "Commands:"
    echo "  list                    - List available configurations"
    echo "  apply [env]            - Apply configuration with environment"
    echo "  validate [env]         - Validate configuration"
    echo "  generate-env [env]     - Generate environment file"
    echo "  diff [env]            - Show diff with current config"
    echo "  restart               - Restart APISIX services"
    echo ""
    echo "Environments:"
    echo "  production            - Production environment"
    echo "  development           - Development environment"
    echo ""
    echo "Examples:"
    echo "  $0 list"
    echo "  $0 apply production"
    echo "  $0 validate development"
    echo "  $0 generate-env production"
    echo "  $0 diff production"
    echo "  $0 restart"
    echo ""
    echo "Configuration Files:"
    echo "  config/config.yaml     - Main APISIX configuration"
    echo "  config/env.production  - Production environment variables"
    echo "  config/env.development - Development environment variables"
}

# Main
case "${1:-help}" in
    list)
        list_configs
        ;;
    apply)
        if [ -z "$2" ]; then
            log_error "Environment required for apply command"
            usage
            exit 1
        fi
        apply_config "$2"
        ;;
    validate)
        if [ -z "$2" ]; then
            log_error "Environment required for validate command"
            usage
            exit 1
        fi
        validate_config "$2"
        ;;
    generate-env)
        if [ -z "$2" ]; then
            log_error "Environment required for generate-env command"
            usage
            exit 1
        fi
        generate_env "$2"
        ;;
    diff)
        if [ -z "$2" ]; then
            log_error "Environment required for diff command"
            usage
            exit 1
        fi
        show_diff "$2"
        ;;
    restart)
        restart_services
        ;;
    help|--help|-h)
        usage
        ;;
    *)
        log_error "Invalid command: $1"
        usage
        exit 1
        ;;
esac
