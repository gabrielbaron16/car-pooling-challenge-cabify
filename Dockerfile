# Stage 1: Build the Go application
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate the server code from the swagger file
RUN swagger generate server -f ./swagger.yml --exclude-main

# Build the Go app
RUN go build -o /carpool ./cmd/car-pooling-server

# Stage 2: Create the final image
FROM alpine:latest

# Get some basic stuff and remove unnecessary apk files
RUN apk --update upgrade && apk add \
  ca-certificates \
  curl \
  tzdata \
  && update-ca-certificates \
  && rm -rf /var/cache/apk/*

# The port your service will listen on
EXPOSE 8080

# We will mark this image with a configurable tag
ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /carpool /carpool

# The command to run
CMD ["/carpool"]