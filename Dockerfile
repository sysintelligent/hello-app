# Start from a minimal base image
FROM golang:1.16-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o hello-app .

# Copy the configuration file
COPY config.json .

# Expose the port on which the application will run
EXPOSE 8080

# Run the Go application
CMD ["./hello-app"]
