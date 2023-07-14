# Start with the official Golang Alpine base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8081

# Set the entry point for the container
CMD ["./main"]
