# Use the official golang image as the base image
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY . .

# Download and install Go module dependencies
RUN go mod download


# Build the Go application as a statically linked binary
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the previous stage
COPY --from=builder /app/main .

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Set the entry point and default arguments
ENTRYPOINT ["./main", "-roads"]

# Define default command arguments
CMD ["a1,a2,a3"]
