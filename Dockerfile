# Stage 1: Install Swagger and generate server code
FROM golang:1.23.4 AS swagger

# Set the Current Working Directory inside the container
WORKDIR /app

# Install Swagger using ADD instruction
ADD https://github.com/go-swagger/go-swagger/releases/download/v0.30.5/swagger_linux_amd64 /usr/local/bin/swagger
RUN chmod +x /usr/local/bin/swagger

# Copy the swagger file
COPY ./swagger.yml /app/swagger.yml

# Generate the server code from the swagger file
RUN /usr/local/bin/swagger generate server -f ./swagger.yml --exclude-main

# Stage 2: Build the Go application
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the generated server code from the swagger stage
COPY --from=swagger /app /app

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o /carpool ./cmd/car-pooling-server

# Stage 3: Create the final image
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