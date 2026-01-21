# --------------------------
# 1️⃣ Build stage
# --------------------------
FROM golang:1.24.3-alpine AS builder

# Install git for Go modules
RUN apk add --no-cache git curl 

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the binary
RUN go build -o auth main.go

# --------------------------
# 2️⃣ Final lightweight image
# --------------------------
FROM alpine:3.18

RUN apk add --no-cache git curl

# Create directory for mounting certificates at runtime
VOLUME /certs

# Copy binary from builder
COPY --from=builder /app/auth /usr/local/bin/auth


# Expose container port
EXPOSE 8080

# Run the Auth service
CMD ["auth"]
