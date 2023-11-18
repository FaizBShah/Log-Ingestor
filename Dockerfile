FROM golang:1.20-alpine AS builder

RUN apk add --no-cache make

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make build

# Expose the port on which the Go application will run
EXPOSE 3000

# Command to run when starting the container.
ENTRYPOINT ["/build/bin/log_ingestor"]