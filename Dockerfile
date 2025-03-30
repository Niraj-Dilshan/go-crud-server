FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

# Create a minimal production image
FROM alpine:latest  

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8888

# Run the binary
CMD ["/app/server"] 