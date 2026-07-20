# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.26.3 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o server .

# Runtime stage
FROM alpine:3.22

WORKDIR /app

# Optional: install CA certificates if your app makes HTTPS requests
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server .

EXPOSE 8081

ENTRYPOINT ["./server"]
