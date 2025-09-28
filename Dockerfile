# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /curaitor ./cmd/curaitor

# Final stage
FROM alpine:latest
WORKDIR /
COPY --from=builder /curaitor /curaitor
# Copy any other necessary files like config files or templates
# For example, if you have a config.yaml you would copy it:
# COPY --from=builder /app/config.yaml /config.yaml

# Expose port if your app listens on one (e.g., 8080)
EXPOSE 9000

ENTRYPOINT ["/curaitor"]
