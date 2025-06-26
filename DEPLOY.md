# Simple Deployment Guide (No Docker)

This guide covers deploying Soko using just a Go binary + Caddy reverse proxy.

## Architecture
```
Internet → Caddy (80/443) → Go Binary (8080) → SQLite File
```

## Prerequisites

### On Your Development Machine
- Go installed
- SSH access to your server

### On Your Server (Ubuntu/Debian)
- Caddy installed
- User account with sudo access

## Server Setup (One-time)

### 1. Install Caddy
```bash
# On your server
sudo apt update
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

### 2. Create Deploy User
```bash
# On your server
sudo adduser deploy
sudo usermod -aG sudo deploy  # Optional: if deploy user needs sudo
```

### 3. Create Application Directories
```bash
# On your server
sudo mkdir -p /opt/soko
sudo mkdir -p /var/lib/soko
sudo chown deploy:deploy /opt/soko /var/lib/soko
```

### 4. Setup Systemd Service
```bash
# Copy the service file to your server
scp soko.service deploy@your-server:/tmp/
ssh deploy@your-server

# On your server
sudo mv /tmp/soko.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable soko
```

## Deployment Process

### 1. Configure Deployment

Edit the `Makefile` and update:
```makefile
SERVER_HOST = your-server-ip  # Your actual server IP
```

Edit `Caddyfile` and update:
```
yourdomain.com {  # Your actual domain
```

Edit `.env/production` and update paths if needed.

### 2. Deploy

From your development machine:
```bash
# Deploy to production
make deploy
```

**What the deploy target does:**
1. Builds your Go binary for Linux (depends on `build-linux` target)
2. Uploads binary and config files to server via SCP
3. Sets up directories and permissions on server
4. Restarts systemd services (soko and caddy)
5. Shows service status

### 3. Verify Deployment

```bash
# Check services are running
ssh deploy@your-server
sudo systemctl status soko
sudo systemctl status caddy

# Check application logs
sudo journalctl -u soko -f

# Test the application
curl https://yourdomain.com/safari/
```

## Manual Deployment (Alternative)

If you prefer to deploy manually:

### 1. Build Binary
```bash
# On your development machine
GOOS=linux GOARCH=amd64 go build -o soko-linux ./cmd/server
```

### 2. Upload Files
```bash
scp soko-linux deploy@your-server:/opt/soko/soko
scp .env/production deploy@your-server:/opt/soko/.env.production
scp Caddyfile deploy@your-server:/tmp/
```

### 3. Setup and Start
```bash
ssh deploy@your-server

# Update Caddy config
sudo mv /tmp/Caddyfile /etc/caddy/Caddyfile
sudo systemctl restart caddy

# Start your app
chmod +x /opt/soko/soko
sudo systemctl restart soko
```

## Database Management

SQLite database is stored at `/var/lib/soko/safari.sqlite` on the server.

### Backup Database
```bash
# On server
cp /var/lib/soko/safari.sqlite /var/lib/soko/safari-backup-$(date +%Y%m%d).sqlite

# Or download to local machine
scp deploy@your-server:/var/lib/soko/safari.sqlite ./safari-backup.sqlite
```

### View Database
```bash
# On server (install sqlite3 if needed)
sudo apt install sqlite3
sqlite3 /var/lib/soko/safari.sqlite
```

## Monitoring & Logs

### View Logs
```bash
# Application logs
sudo journalctl -u soko -f

# Caddy logs
sudo journalctl -u caddy -f

# Or view Caddy access logs
sudo tail -f /var/log/caddy/soko.log
```

### Service Management
```bash
# Restart application
sudo systemctl restart soko

# Restart Caddy
sudo systemctl restart caddy

# Check status
sudo systemctl status soko caddy
```

## Updating Your App

Just run the deploy command again:
```bash
make deploy
```

It will:
1. Build new binary for Linux
2. Upload and replace old binary
3. Restart service with zero downtime (systemd handles this)

## Troubleshooting

### App won't start
```bash
# Check logs
sudo journalctl -u soko --no-pager

# Check if port is available
sudo netstat -tulpn | grep :8080

# Test binary manually
cd /opt/soko
GO_ENV=production ./soko
```

### SSL certificate issues
```bash
# Check Caddy config
sudo caddy validate --config /etc/caddy/Caddyfile

# Check Caddy logs
sudo journalctl -u caddy --no-pager
```

### Database permission issues
```bash
# Fix ownership
sudo chown deploy:deploy /var/lib/soko/safari.sqlite
sudo chmod 644 /var/lib/soko/safari.sqlite
```

## Performance Notes

This simple setup can easily handle:
- **Thousands of concurrent users**
- **Millions of database records** (SQLite is very capable)
- **High availability** with proper server setup

For scaling later, consider:
- Load balancer + multiple app instances
- Database replication
- CDN for static assets

But you're good for a LONG time with this simple setup! 🚀