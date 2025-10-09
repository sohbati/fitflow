# CoachAidApp Admin Gateway Configuration

This file contains the default configuration for the APISIX admin gateway.

## Configuration Overview

The gateway is configured with:
- **Node Listen Port**: 9080 (main gateway port)
- **Admin API Port**: 9180 (management interface)
- **ETCD Port**: 23791 (configuration storage)
- **Admin Key**: admin-key (for API authentication)

## Default Routes

The gateway comes pre-configured with basic routes for the CoachAidApp:

### 1. Frontend Route
- **Path**: `/`
- **Target**: Angular frontend (client)
- **Port**: 4200

### 2. Backend API Route
- **Path**: `/api/*`
- **Target**: Java backend (server)
- **Port**: 8080

### 3. Admin Gateway Route
- **Path**: `/admin/*`
- **Target**: Admin gateway management
- **Port**: 9180

## Environment Variables

You can override default settings using environment variables:

```bash
export GATEWAY_PORT=9080
export ADMIN_PORT=9180
export ETCD_PORT=23791
export ADMIN_KEY=admin-key
```

## Service Configuration

### Frontend Service (Angular)
```yaml
service:
  name: "coach-aid-frontend"
  upstream:
    nodes:
      "127.0.0.1:4200": 1
```

### Backend Service (Java)
```yaml
service:
  name: "coach-aid-backend"
  upstream:
    nodes:
      "127.0.0.1:8080": 1
```

## Plugin Configuration

### Rate Limiting
```yaml
plugins:
  limit-req:
    rate: 100
    burst: 200
    key: "remote_addr"
```

### CORS
```yaml
plugins:
  cors:
    allow_origins: "*"
    allow_methods: "GET,POST,PUT,DELETE,OPTIONS"
    allow_headers: "Content-Type,Authorization"
```

### Authentication
```yaml
plugins:
  jwt-auth:
    key: "coach-aid-jwt"
    secret: "your-secret-key"
```

## Health Checks

The gateway includes health check endpoints:
- **Gateway Health**: `GET /health`
- **Admin Health**: `GET /admin/health`

## Monitoring

Basic monitoring is configured:
- **Metrics**: Available at `/metrics`
- **Logs**: Stored in `apisix/logs/`
- **Status**: Available at `/status`

## Security

Default security settings:
- **Admin Access**: Restricted to localhost
- **API Keys**: Required for admin operations
- **Rate Limiting**: Enabled by default
- **CORS**: Configured for development

## Customization

To customize the configuration:

1. **Edit Config**: Modify `apisix/conf/config.yaml`
2. **Restart Gateway**: Run `./admin-gateway.sh restart`
3. **Verify Changes**: Check status with `./admin-gateway.sh status`

## Development vs Production

### Development
- CORS enabled for all origins
- Admin access from localhost only
- Basic rate limiting
- Debug logging enabled

### Production
- Restricted CORS origins
- Strong admin authentication
- Aggressive rate limiting
- Error logging only
- SSL/TLS termination
