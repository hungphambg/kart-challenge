# Stage 1: Builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Stage 2: Runner
FROM alpine:latest

# Create a non-root user
RUN adduser -D appuser
USER appuser

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Copy migration files
COPY --from=builder /app/migrations ./migrations

# Expose the port the app runs on
EXPOSE 1323

# Command to run the executable
CMD ["./main"]
