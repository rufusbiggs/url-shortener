# Build stage - use golang:1.23 (includes full glibc version support)
FROM golang:1.23 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application for Linux target (cross-compiling)
RUN GOOS=linux GOARCH=amd64 go build -o url-shortener .

# Runtime stage - use a full Ubuntu 22.04 base image for glibc compatibility
FROM ubuntu:22.04

# Install required dependencies, including glibc
RUN apt-get update && apt-get install -y libc6

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/url-shortener .

# Copy the init.sql file into the PostgreSQL container's initialization directory
COPY ./init.sql /docker-entrypoint-initdb.d/

# Set executable permissions
RUN chmod +x /root/url-shortener

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./url-shortener"]
