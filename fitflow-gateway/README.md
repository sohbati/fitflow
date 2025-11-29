# FitFlow Gateway (Nginx)

Pure Nginx-based API gateway for FitFlow services.

## Prerequisites

- Nginx installed on your system
  - macOS: `brew install nginx`
  - Linux: `apt-get install nginx` or `yum install nginx`

## Project layout

- `config/nginx.conf` – main Nginx configuration file
- `config/conf.d/` – additional configuration files (upstreams, routes, etc.)
- `logs/` – Nginx access and error logs (gitignored)
- `run-fitflow-gateway.sh` – helper script at repository root to manage Nginx

## Usage

From the repository root:

```bash
./run-fitflow-gateway.sh start    # start Nginx
./run-fitflow-gateway.sh stop     # stop Nginx
./run-fitflow-gateway.sh restart  # restart Nginx
./run-fitflow-gateway.sh reload   # reload configuration
./run-fitflow-gateway.sh status   # check Nginx status
./run-fitflow-gateway.sh logs     # tail Nginx logs
```

Once running:

- Gateway URL: http://localhost:8080
- Admin/Status: http://localhost:8080/status (if enabled)

Edit `config/nginx.conf` or files in `config/conf.d/` to add new upstream services, routes, and configurations. After editing, run `./run-fitflow-gateway.sh reload` to apply changes without downtime.

