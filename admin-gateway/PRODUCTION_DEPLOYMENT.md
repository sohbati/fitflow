# CoachAidApp APISIX Production Deployment Guide

This guide covers deploying the CoachAidApp APISIX gateway in production environments using Ubuntu or Alpine Linux.

## 🚀 Quick Start

### Native Linux Installation
```bash
# Run as root
sudo ./admin-gateway-linux.sh ubuntu production

# Check status
systemctl status apisix etcd
```

## 📋 Prerequisites

### System Requirements
- **OS**: Ubuntu 20.04+ or Alpine Linux 3.15+
- **RAM**: Minimum 2GB, Recommended 4GB+
- **Disk**: Minimum 10GB free space
- **CPU**: 2+ cores recommended
- **Network**: Internet connection for package downloads

### Required Ports
- **9080**: APISIX Gateway (HTTP)
- **9180**: APISIX Admin API
- **2379**: ETCD (internal)

## 🔧 Configuration

### Environment Variables
Create `/etc/apisix.env` with your production settings:

```bash
# Production Environment Configuration
APISIX_ADMIN_KEY=your-secure-admin-key-here
ETCD_HOST=http://127.0.0.1:2379
APISIX_NODE_PORT=9080
APISIX_ALLOW_ADMIN=your-admin-ip-address
NGINX_WORKER_PROCESSES=auto
```

### Configuration Files
- **Main Config**: `config/config.yaml` - Single APISIX configuration
- **Environment Files**: 
  - `config/env.production` - Production environment variables
  - `config/env.development` - Development environment variables

## 🐧 Native Linux Deployment

### Ubuntu/Debian
```bash
# Run deployment script
sudo ./admin-gateway-linux.sh ubuntu production

# Check service status
systemctl status apisix etcd

# View logs
journalctl -u apisix -f
journalctl -u etcd -f
```

### Alpine Linux
```bash
# Run deployment script
sudo ./admin-gateway-linux.sh alpine production

# Check service status
rc-status
service apisix status
service etcd status
```

## 🔒 Security Configuration

### Firewall Setup
```bash
# Ubuntu/Debian (UFW)
ufw allow 9080/tcp
ufw allow 9180/tcp
ufw allow 9000/tcp  # Dashboard
ufw deny 2379/tcp   # ETCD (internal only)

# Alpine Linux (iptables)
iptables -A INPUT -p tcp --dport 9080 -j ACCEPT
iptables -A INPUT -p tcp --dport 9180 -j ACCEPT
iptables -A INPUT -p tcp --dport 9000 -j ACCEPT
iptables -A INPUT -p tcp --dport 2379 -j DROP
```

### SSL/TLS Configuration
```yaml
# In config.yaml
apisix:
  ssl:
    ssl_cert: /path/to/cert.pem
    ssl_cert_key: /path/to/key.pem
    listen_http2: true
```

### Admin Access Restriction
```yaml
# Restrict admin access to specific IPs
apisix:
  allow_admin:
    - 192.168.1.100
    - 10.0.0.50
```

## 📊 Monitoring & Logging

### Health Checks
```bash
# Check APISIX health
curl http://localhost:9080/health

# Check ETCD health
curl http://localhost:2379/health
```

### Log Management
```bash
# View APISIX logs
tail -f /var/log/apisix/error.log
tail -f /var/log/apisix/access.log

# Docker logs
docker-compose logs -f apisix
```

### Prometheus Metrics
```bash
# Enable Prometheus plugin
curl -X POST http://localhost:9180/apisix/admin/routes/1 \
  -H "X-API-KEY: your-admin-key" \
  -H "Content-Type: application/json" \
  -d '{
    "uri": "/metrics",
    "plugins": {
      "prometheus": {}
    }
  }'
```

## 🔄 Service Management

### Systemd Commands
```bash
# Start services
systemctl start etcd apisix

# Stop services
systemctl stop apisix etcd

# Restart services
systemctl restart apisix etcd

# Enable auto-start
systemctl enable etcd apisix

# Check status
systemctl status apisix etcd
```

### Service Management Commands
```bash
# Start specific service
systemctl start apisix

# Restart service
systemctl restart apisix

# Enable auto-start
systemctl enable apisix etcd
```

## 🚨 Troubleshooting

### Common Issues

#### 1. APISIX Won't Start
```bash
# Check configuration
/opt/apisix/bin/apisix test

# Check logs
journalctl -u apisix -n 50

# Verify ETCD connection
curl http://localhost:2379/health
```

#### 2. ETCD Connection Issues
```bash
# Check ETCD status
systemctl status etcd

# Check ETCD logs
journalctl -u etcd -n 50

# Restart ETCD
systemctl restart etcd
```

#### 3. Port Conflicts
```bash
# Check port usage
netstat -tlnp | grep :9080
netstat -tlnp | grep :9180

# Kill conflicting processes
sudo kill -9 <PID>
```

### Performance Tuning

#### Nginx Configuration
```yaml
nginx_config:
  worker_processes: auto
  worker_rlimit_nofile: 65535
  events:
    worker_connections: 4096
    use: epoll
    multi_accept: on
```

#### ETCD Tuning
```bash
# Increase ETCD limits
echo "etcd soft nofile 65536" >> /etc/security/limits.conf
echo "etcd hard nofile 65536" >> /etc/security/limits.conf
```

## 📈 Scaling

### Horizontal Scaling
```bash
# Multiple APISIX instances
# Deploy APISIX on multiple servers
# Use nginx or HAProxy as load balancer in front of APISIX instances
```

### Vertical Scaling
```yaml
# Increase worker processes
nginx_config:
  worker_processes: 8
  worker_rlimit_nofile: 65535
```

## 🔄 Updates & Maintenance

### Updating APISIX
```bash
# Stop services
systemctl stop apisix

# Backup configuration
cp /opt/apisix/conf/config.yaml /backup/

# Update APISIX
wget https://github.com/apache/apisix/releases/download/3.9.0/apisix-3.9.0-src.tar.gz
tar -xzf apisix-3.9.0-src.tar.gz
cd apisix-3.9.0
make deps

# Restore configuration
cp /backup/config.yaml /opt/apisix/conf/

# Start services
systemctl start apisix
```

### Backup Strategy
```bash
# Backup ETCD data
etcdctl snapshot save /backup/etcd-backup.db

# Backup configuration
tar -czf /backup/apisix-config-$(date +%Y%m%d).tar.gz /opt/apisix/conf/
```

## 📞 Support

For issues related to:
- **APISIX**: Check [APISIX Documentation](https://apisix.apache.org/docs/)
- **ETCD**: Check [ETCD Documentation](https://etcd.io/docs/)
- **CoachAidApp**: Check project documentation

## 🎯 Next Steps

After successful deployment:
1. Configure routes for your CoachAidApp services
2. Set up SSL certificates
3. Configure monitoring and alerting
4. Set up backup procedures
5. Test failover scenarios
