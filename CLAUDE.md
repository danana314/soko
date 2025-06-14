# Claude Code Instructions

This file contains instructions and context for Claude Code to help with development tasks.

## Project Overview
This is a trip expense sharing application called "Soko" built with Go, SQLite, and HTMX. Users can create trips, add participants, manage schedules, and track shared expenses.

## Architecture
- **Backend**: Go with standard library HTTP server
- **Database**: SQLite with schema in `internal/store/schema.sql`
- **Frontend**: HTMX + Alpine.js with server-side templates
- **Styling**: Tailwind CSS + DaisyUI

## Key Files
- `cmd/handlers.go` - HTTP handlers for all routes
- `internal/models/trip.go` - Data models (Trip, User, Expense, etc.)
- `internal/store/data.go` - Database operations
- `web/templates/trip.tmpl` - Main UI template
- `internal/funcs/funcs.go` - Template helper functions

## Development Commands
- **Build**: `go build -o tmp/main cmd/*.go`
- **Run**: `./tmp/main` (serves on localhost:8080)
- **Test**: `go test ./...`
- **Database**: SQLite file at `db.sqlite` (DO NOT delete - contains user data)

## Database Schema
- `trips`: Basic trip info (name, dates)
- `users`: Trip participants
- `schedule`: User availability by date
- `expenses`: Shared expenses with JSON participant lists

## Code Conventions
- Use existing patterns from the codebase
- Follow Go naming conventions
- Template functions are in `internal/funcs/funcs.go`
- Database operations use prepared statements
- JSON is used for complex data in SQLite (like expense participants)

## Common Tasks
- **Adding routes**: Update `cmd/handlers.go` routes() function
- **New templates**: Add to `web/templates/` and reference in handlers
- **Database changes**: Update `internal/store/schema.sql` and data access functions
- **Template helpers**: Add to `internal/funcs/funcs.go` and TemplateFuncs map

## Important Notes
- The `db.sqlite` file contains user data - never delete it
- Expense amounts are stored as strings for precise decimal handling
- Participants in expenses are stored as JSON arrays
- HTMX is used for dynamic updates without page reloads
- Alpine.js handles client-side interactivity (tabs, etc.)