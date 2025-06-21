# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o soko cmd/*.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Create non-root user
RUN addgroup -g 1001 -S soko && \
    adduser -S soko -u 1001 -G soko

# Copy the binary from builder stage
COPY --from=builder /app/soko .

# Create directory for database and set permissions
RUN mkdir -p /data && chown soko:soko /data

# Switch to non-root user
USER soko

# Expose port
EXPOSE 8080

# Set environment variables
ENV DB_PATH=/data/db.sqlite
ENV PORT=8080

# Run the application
CMD ["./soko"]