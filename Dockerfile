# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

# Install git (needed for go get) and ca-certificates
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Set necessary environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire app
COPY . .

# Build the binary with additional security flags
RUN go build -ldflags="-w -s" -o server ./cmd/server


# ---------- Stage 2: Runtime ----------
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy binary from builder
COPY --from=builder /app/server /server

# Run as non-root
USER appuser

# Expose the port your app runs on (adjust if different)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/server"]