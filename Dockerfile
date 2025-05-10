# Use the official Go runtime as the base image
FROM golang:1.23-alpine AS builder

# Install git and other necessary tools
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Set Go proxy to direct
ENV GOPROXY=direct

ENV GOMODCACHE="/app/go/pkg/mod"

ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GOARCH=amd64

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy && go mod verify

RUN go mod download

# Copy the project source code
COPY . .

# Debug: List files in the working directory
RUN ls -l /app

# Build the Gin application
RUN --mount=type=cache,target=/app/go/pkg/mod go build -o /binary ./src/cmd/main/main.go

# Use a minimal Alpine image as the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /

# Copy the compiled Gin binary from the builder stage
COPY --from=builder /binary /app

# Expose the port your Gin app listens on (e.g., 8081)
EXPOSE 8083

# Set the command to run the Gin application
CMD ["/app"]