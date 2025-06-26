# Claude Code Instructions

This file contains instructions and context for Claude Code to help with development tasks.

## Project Overview
This is a monorepo called "Soko" containing multiple Go applications served via subdirectory routing. The main app is "Safari" - a trip expense sharing application built with Go, SQLite, and HTMX.

## Architecture
- **Monorepo Structure**: Single Go server with subdirectory routing (`/safari/`, future apps at `/app2/`, etc.)
- **Backend**: Go with standard library HTTP server
- **Database**: Separate SQLite databases per app (`data/safari.sqlite`)
- **Frontend**: HTMX + Alpine.js with server-side templates
- **Styling**: Tailwind CSS with custom Sardinian color palette
- **Module**: `1008001/soko`

## Key Files
- `cmd/server/main.go` - Main server with subdirectory routing
- `cmd/apps/safari/handlers.go` - Safari app HTTP handlers
- `cmd/apps/safari/routes.go` - Safari app route setup
- `internal/safari/models/trip.go` - Safari app data models
- `internal/safari/store/data.go` - Safari app database operations
- `web/safari/templates/trip.tmpl` - Safari app UI templates
- `internal/safari/funcs/funcs.go` - Safari app template helpers

## Development Commands
- **Build**: `make build` (outputs to `/tmp/bin/soko`)
- **Run**: `make run` (builds and runs server on localhost:8080)
- **Run with hot reload**: `make run/live`
- **Test**: `go test ./...`
- **Build for Linux**: `make build-linux` (for deployment)
- **Deploy**: `make deploy` (runs deploy.sh script)
- **Database**: SQLite files in `data/` directory (DO NOT delete - contains user data)

## Database Schema (Safari App)
- `trips`: Basic trip info (name, dates)
- `users`: Trip participants
- `schedule`: User availability by date
- `expenses`: Shared expenses with JSON participant lists

## Code Conventions
- Use existing patterns from the codebase
- Follow Go naming conventions
- Safari app routes are prefixed with `/safari/` (e.g., `/safari/t/{tripId}`)
- Template functions are in `internal/safari/funcs/funcs.go`
- Database operations use prepared statements
- JSON is used for complex data in SQLite (like expense participants)
- Use Sardinian color palette: sardinia-orange (apricot #f6b55f), sardinia-turquoise, sardinia-terracotta, sardinia-cream, etc.

## Common Tasks
- **Adding Safari routes**: Update `cmd/apps/safari/handlers.go` routes() function
- **New Safari templates**: Add to `web/safari/templates/` and reference in handlers
- **Database changes**: Update `internal/safari/store/schema.sql` and data access functions
- **Template helpers**: Add to `internal/safari/funcs/funcs.go` and TemplateFuncs map
- **Adding new apps**: Create new directory under `cmd/apps/` and `internal/`, register in main server

## Important Notes
- Database files in `data/` contain user data - never delete them
- Environment variables are loaded from `.env.development` or `.env.production` via godotenv
- Safari app is accessible at `http://localhost:8080/safari/`
- Root path `/` redirects to Safari app
- Expense amounts are stored as strings for precise decimal handling
- Participants in expenses are stored as JSON arrays
- HTMX is used for dynamic updates without page reloads
- Alpine.js handles client-side interactivity (tabs, etc.)
- Uses custom Sardinian color theme inspired by Mediterranean aesthetics
- No DaisyUI - pure Tailwind CSS with custom components

## Adding New Apps
To add a new app to the monorepo:
1. Create `cmd/apps/newapp/` with handlers and routes
2. Create `internal/newapp/` with business logic
3. Create `web/newapp/` with templates and static files
4. Register the app in `cmd/server/main.go` with subdirectory routing
5. Add environment variables for the new app's database

## Deployment (Simple Binary + Caddy)
The app uses a simple deployment approach:
- **Architecture**: Internet → Caddy (80/443) → Go Binary (8080) → SQLite File
- **Deploy command**: `make deploy` builds and deploys to server
- **Configuration**: Update `SERVER_HOST` in Makefile with your server IP
- **Reverse proxy**: Caddy handles SSL, compression, security headers
- **Process manager**: systemd manages the Go binary as a service
- **Database**: SQLite file at `/var/lib/soko/safari.sqlite` on server
- **Logs**: `sudo journalctl -u soko -f`

See `DEPLOY.md` for complete setup instructions.
