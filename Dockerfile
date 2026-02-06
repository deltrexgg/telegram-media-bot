# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bot ./cmd/bot

# ---- Runtime stage ----
FROM alpine:latest

WORKDIR /app

# Required for SQLite + certificates
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bot /app/bot

# DB will live here (mounted volume)
VOLUME ["/data"]

ENV DB_PATH=/data/database.db

CMD ["/app/bot"]
