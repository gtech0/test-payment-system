# ---------- Stage 1: Build ----------
FROM golang:1.24 AS builder

WORKDIR /app

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary for linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/app .

# ---------- Stage 2: Runtime ----------
FROM alpine:3.20

# for TLS support if your app connects over TLS
RUN apk --no-cache add ca-certificates

# non-root user (recommended)
RUN addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=builder /app/app /app/app
RUN chown app:app /app/app
USER app

# Expose your app port (optional documentation)
EXPOSE 8001

CMD ["/app/app"]
