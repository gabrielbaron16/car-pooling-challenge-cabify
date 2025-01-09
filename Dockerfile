# Use the official Golang image as the build stage
FROM golang:1.23.4-alpine AS builder

# Install dependencies
RUN apk --update upgrade && apk add --no-cache \
  ca-certificates \
  curl \
  tzdata

# Set the working directory
WORKDIR /app

# Copy Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -a -o target/bin/carpool ./cmd/car-pooling-server/main.go

# Use a minimal base image for the final image
FROM alpine:latest

# Install necessary packages
RUN apk --update upgrade && apk add --no-cache \
  ca-certificates \
  tzdata

# Copy the built binary from the builder stage
COPY --from=builder /app/target/bin/carpool /carpool

# Expose the port your service will listen on
EXPOSE 8080

# Set build arguments and labels
ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

# Set the command to run the application
CMD ["/carpool"]