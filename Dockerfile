# ---------- Builder ----------
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Needed for sqlite
RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the bot
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -o bot ./cmd/bot

# ---------- Runtime ----------
FROM alpine:latest

WORKDIR /app

# SQLite runtime dependency
RUN apk add --no-cache sqlite-libs ca-certificates

COPY --from=builder /app/bot /app/bot

# Directory where DB will be mounted
RUN mkdir -p /data

ENV DB_PATH=/data/database.db

CMD ["/app/bot"]

