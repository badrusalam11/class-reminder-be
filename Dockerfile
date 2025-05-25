# Build stage
FROM golang:1.22.2 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code (including config.json)
COPY . .

# Build the Go app with output binary name
RUN go build -o class-reminder-be .

# Final stage
FROM debian:bookworm-slim

# Set working directory
WORKDIR /app

# Copy binary and config.json from builder stage
COPY --from=builder /app/class-reminder-be .
COPY --from=builder /app/config.json .
COPY --from=builder /app/.env .

# Expose application port
EXPOSE 9090

# Run the app
CMD ["./class-reminder-be"]
