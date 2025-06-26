# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the cmd/server application
.PHONY: build
build:
	go build -o=/tmp/bin/soko ./cmd/server

## run: run the cmd/server application
.PHONY: run
run: build
	/tmp/bin/soko

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/soko" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# ==================================================================================== #
# DEPLOYMENT
# ==================================================================================== #

# Deployment configuration - update these for your server
SERVER_USER = deploy
SERVER_HOST = your-server-ip
APP_NAME = soko
REMOTE_DIR = /opt/soko

## build-linux: build binary for Linux deployment
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux ./cmd/server

## deploy: deploy application to production server
.PHONY: deploy
deploy: build-linux
	@echo "🚀 Deploying Soko to production..."
	@echo "⬆️  Uploading files to server..."
	scp $(APP_NAME)-linux $(SERVER_USER)@$(SERVER_HOST):$(REMOTE_DIR)/$(APP_NAME)
	scp .env.production $(SERVER_USER)@$(SERVER_HOST):$(REMOTE_DIR)/.env.production
	scp Caddyfile $(SERVER_USER)@$(SERVER_HOST):/tmp/Caddyfile
	@echo "🔄 Restarting services on server..."
	ssh $(SERVER_USER)@$(SERVER_HOST) '\
		sudo mkdir -p /var/lib/soko && \
		sudo chown deploy:deploy /var/lib/soko && \
		sudo chmod +x $(REMOTE_DIR)/$(APP_NAME) && \
		sudo mv /tmp/Caddyfile /etc/caddy/Caddyfile && \
		sudo systemctl restart $(APP_NAME) && \
		sudo systemctl restart caddy && \
		sudo systemctl status $(APP_NAME) --no-pager && \
		sudo systemctl status caddy --no-pager'
	rm $(APP_NAME)-linux
	@echo "✅ Deployment complete!"
	@echo "🌐 Your app should be available at https://1008001.xyz/safari/"
