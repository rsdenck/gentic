# Multi-stage build for Infrastructure AI Agents
# Stage 1: Build
FROM golang:1.26.7-alpine AS builder

WORKDIR /app

# Copy go mod and sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY .env .env

# Build with CGO disabled for static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server/

# Stage 2: Minimal runtime
FROM alpine:3.19

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .
COPY --from=builder /app/.env .env

# Expose the application port
EXPOSE 8080

# Set environment variables defaults
ENV OLLAMA_SERVER_ADDRESS=https://api.denck.sbs
ENV OLLAMA_MODEL=muse-glimmer:30b-q4_K_M-dflash
ENV OLLAMA_TIMEOUT=300
ENV APP_PORT=8080
ENV APP_ENV=production

# Entrypoint
ENTRYPOINT ["/server"]

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1