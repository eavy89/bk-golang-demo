# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

# Install git, ca-certificates, and build dependencies (gcc, musl-dev)
RUN apk add --no-cache git ca-certificates gcc musl-dev && update-ca-certificates

# Set necessary environment variables
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire app
COPY . .

# Build the binary with additional security flags
RUN go build -ldflags="-w -s" -o server ./src/app


# ---------- Stage 2: Runtime ----------
FROM alpine:latest

# Install SQLite dependencies
RUN apk add --no-cache libc6-compat

# Create directory structure with correct permissions
RUN mkdir -p /app/data && \
    addgroup -S -g 1000 appgroup && \
    adduser -S -u 1000 -G appgroup appuser && \
    chown -R 1000:1000 /app

WORKDIR /app

# Copy binary and envirnments from builder
COPY --from=builder /app/server /app/server
COPY .env /app/.env
RUN chown appuser:appgroup /app/server /app/.env

# Run as non-root
#USER 1000

# Expose the port your app runs on (adjust if different)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/app/server"]