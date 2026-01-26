# Stage 1: Builder
FROM golang:1.22-alpine AS builder

WORKDIR /app/consumer

# Copy go.mod and go.sum files and download dependencies
COPY pipeline/consumer/go.mod pipeline/consumer/go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY pipeline/consumer/. .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runner
FROM alpine:latest

# Create a non-root user
RUN adduser -D appuser
USER appuser

WORKDIR /app/consumer

# Copy the built binary from the builder stage
COPY --from=builder /app/consumer/main .

# Command to run the executable
ENTRYPOINT ["./main"]
