# Start with the official Golang image as a base
FROM golang:1.22-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache module dependencies
RUN go mod download

# Create bin directory for the binary
RUN mkdir -p bin

# Copy the source code to the working directory
COPY ./src ./src

# Build the application
RUN go build -o bin ./src

# Debugging step to check if the binary exists
RUN ls -l bin

# Expose application port
EXPOSE 8080

# Command to run the application
CMD ["./bin"]
