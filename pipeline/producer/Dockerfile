# Stage 1: Builder
FROM golang:1.22-alpine AS builder

WORKDIR /app/producer

# Copy go.mod and go.sum files and download dependencies
COPY pipeline/producer/go.mod pipeline/producer/go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY pipeline/producer/. .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runner
FROM alpine:latest

# Create a non-root user
RUN adduser -D appuser
USER appuser

WORKDIR /app/producer

# Copy the built binary from the builder stage
COPY --from=builder /app/producer/main .

# Command to run the executable
# Expects filename as argument
ENTRYPOINT ["./main"]
