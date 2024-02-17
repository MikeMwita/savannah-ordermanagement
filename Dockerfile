# Stage 1: Build Environment
FROM golang:1.17-alpine as build-env

# Set environment variables
ENV APP_NAME savannah
ENV CMD_PATH cmd/api/

# Set the working directory
WORKDIR /go/src/$APP_NAME

# Copy the source code into the container's workspace
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -v -o /go/bin/$APP_NAME $CMD_PATH

# Stage 2: Final Image
FROM alpine:3.14

# Install Go
RUN apk add --no-cache go

# Copy the built binary from the previous stage
COPY --from=build-env /go/bin/gophercon-backend /go/bin/savannah

# Set environment variables
ENV GO111MODULE=on

# Set the entry point
CMD ["/go/bin/gophercon-backend"]