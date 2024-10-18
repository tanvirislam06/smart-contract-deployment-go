# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Create directories inside the container
RUN mkdir -p /app/contracts /app/scripts /app/bindings /app/build

# Copy the source code into the container
COPY contracts/ /app/contracts/
COPY scripts/ /app/scripts/
COPY bindings/ /app/bindings/
COPY entrypoint.sh /app/
COPY Dockerfile /app/

# Build the Go app
RUN go build -o /app/deploy /app/scripts/Deploy.go

# Ensure the entrypoint script is executable
RUN chmod +x /app/entrypoint.sh

# Command to run the entrypoint script
ENTRYPOINT ["/app/entrypoint.sh"]