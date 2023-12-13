# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files for dependency resolution
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .


#Running the producer and the consumer
CMD ["go", "run", "producer.go"]


CMD ["go", "run", "consumer.go"]



