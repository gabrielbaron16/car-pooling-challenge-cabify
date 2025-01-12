FROM alpine:latest

WORKDIR /app

COPY . .

# Get some basic stuff and remove innecessary apk files
RUN apk --update upgrade && apk add \
  ca-certificates \
  curl \
  tzdata \
  && update-ca-certificates \
  && rm -rf /var/cache/apk/*

# We will mark this image with a configurable tag
ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

# Copy the service binary
COPY target/bin/carpool /carpool

# The command to run
CMD ["/carpool", "--host", "0.0.0.0"]

# The port your service will listen on
EXPOSE 9091