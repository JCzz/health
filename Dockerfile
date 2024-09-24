# Use the official Golang image with Go 1.22 as the builder
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies (caching layer)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o health-server

# Use a minimal base image to run the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/health-server .

# Expose the server port
EXPOSE 8080

# Set the HEALTH_STATUS environment variable (can be overridden)
# ENV HEALTH_STATUS="Healthy"

# Run the web server
CMD ["./health-server"]
