# FitFlow Admin Gateway

This directory contains the APISIX admin gateway setup for the FitFlow project.

## Overview

The admin gateway provides:
- **API Gateway**: Routes requests to backend services
- **Admin API**: Management interface for configuring routes, plugins, and services
- **Load Balancing**: Distributes traffic across multiple backend instances
- **Authentication**: Handles API authentication and authorization
- **Rate Limiting**: Protects backend services from abuse

## Scripts Overview

This project includes three specialized scripts for different purposes:

### 1. **`admin-gateway-macos.sh`** - macOS Development
- **Purpose**: Local development on macOS
- **Features**: Homebrew dependency management, local APISIX instance
- **Usage**: `./admin-gateway-macos.sh [install|start|stop|restart|status]`

### 2. **`admin-gateway-linux.sh`** - Linux Production Deployment
- **Purpose**: Production deployment on Ubuntu/Alpine Linux servers
- **Features**: Full system installation, systemd services, production configuration
- **Usage**: `sudo ./admin-gateway-linux.sh [ubuntu|alpine] [production|development]`

### 3. **`config-manager.sh`** - Configuration Management
- **Purpose**: Manage configurations across environments
- **Features**: Apply configs, validate, generate environment files, show diffs
- **Usage**: `./config-manager.sh [list|apply|validate|generate-env|diff|restart]`

## Quick Start

### For macOS Development
```bash
# Install dependencies
./admin-gateway-macos.sh install

# Start the gateway
./admin-gateway-macos.sh start

# Check status
./admin-gateway-macos.sh status

# Stop the gateway
./admin-gateway-macos.sh stop

# Restart the gateway
./admin-gateway-macos.sh restart
```

### For Linux Production Deployment
```bash
# Deploy on Ubuntu
sudo ./admin-gateway-linux.sh ubuntu production

# Deploy on Alpine Linux
sudo ./admin-gateway-linux.sh alpine production
```

## Configuration

### Single Configuration File
The gateway uses a single, flexible configuration file that works for both development and production:

- **Main Config**: `config/config.yaml` - Single APISIX configuration with environment variables
- **Environment Files**: 
  - `config/env.production` - Production environment variables
  - `config/env.development` - Development environment variables
  - `config/env.example` - Template for custom environments

### Default Settings
- **Gateway Port**: 9080
- **Admin API Port**: 9180
- **ETCD Port**: 2379 (production) / 23791 (development)
- **Admin Key**: Configurable via environment variables

### Gateway URLs
- **Gateway**: http://localhost:9080
- **Admin API**: http://localhost:9180

## Dependencies

Before running the gateway, ensure you have:
- **Git**: For cloning APISIX repository
- **etcd**: For configuration storage
- **Lua**: For APISIX runtime

### Installation

#### macOS (using Homebrew)
```bash
brew install git etcd lua
```

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install git etcd lua5.1
```

#### CentOS/RHEL
```bash
sudo yum install git etcd lua
```

## Usage Examples

### 1. Start the Gateway
```bash
./admin-gateway.sh start
```

### 2. Create a Route
```bash
curl -X POST http://localhost:9180/apisix/admin/routes/1 \
  -H "X-API-KEY: admin-key" \
  -H "Content-Type: application/json" \
  -d '{
    "uri": "/api/*",
    "upstream": {
      "type": "roundrobin",
      "nodes": {
        "127.0.0.1:8080": 1
      }
    }
  }'
```

### 3. Test the Route
```bash
curl http://localhost:9080/api/health
```

## Project Integration

This gateway is designed to work with the CoachAidApp project structure:

```
CoachAidApp/
├── admin-gateway/          # This directory
│   ├── admin-gateway.sh    # Gateway management script
│   ├── apisix/             # APISIX installation (auto-created)
│   └── etcd-data/          # ETCD data directory (auto-created)
├── client/                 # Angular frontend
├── server/                 # Java backend
└── ...
```

## Configuration Files

The gateway automatically creates configuration files:
- `apisix/conf/config.yaml` - Main APISIX configuration
- `etcd-data/` - ETCD data directory

## Troubleshooting

### Common Issues

1. **Port Already in Use**
   ```bash
   # Check what's using the port
   lsof -i :9080
   lsof -i :9180
   lsof -i :23791
   ```

2. **Permission Denied**
   ```bash
   # Make script executable
   chmod +x admin-gateway.sh
   ```

3. **Missing Dependencies**
   ```bash
   # Check if required tools are installed
   which git etcd lua
   ```

### Logs

- **APISIX Logs**: Check `apisix/logs/` directory
- **ETCD Logs**: Check system logs or run with verbose output

## Development

For development purposes, you can:

1. **Modify Configuration**: Edit `apisix/conf/config.yaml`
2. **Add Plugins**: Configure additional APISIX plugins
3. **Custom Routes**: Set up routes for different services

## Production Considerations

For production deployment:
- Use proper SSL certificates
- Configure firewall rules
- Set up monitoring and logging
- Use external etcd cluster
- Configure proper admin authentication

## Support

For issues related to:
- **APISIX**: Check [APISIX Documentation](https://apisix.apache.org/docs/)
- **ETCD**: Check [ETCD Documentation](https://etcd.io/docs/)
- **CoachAidApp**: Check project documentation
