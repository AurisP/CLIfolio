FROM golang:1.23-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o cli-app .

# Use a minimal base image to keep the image size small
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/cli-app .

# Set the entrypoint to run the CLI app
ENTRYPOINT ["./cli-app"]
