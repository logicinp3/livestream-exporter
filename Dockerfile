# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o livestream-exporter .

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/livestream-exporter .
ARG APP_PORT=8080
EXPOSE ${APP_PORT}
CMD ["/app/livestream-exporter"]
