# ---- Build Stage ----
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o webssh ./cmd/webssh

# ---- Runtime Stage ----
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/webssh .

RUN mkdir -p /app/data && chmod 755 /app/data

# Data directory will be created at runtime if needed
VOLUME ["/app/data"]

# Environment variables (override with -e):
# PORT  - HTTP port (default: 8888)
# AUTH  - Enable auth (default: false)
# STORE - Enable storage (default: false)
ENV PORT=8888 AUTH=false STORE=false

EXPOSE 8888

CMD ["sh", "-c", "./webssh --port $PORT --auth $AUTH --store $STORE"]
