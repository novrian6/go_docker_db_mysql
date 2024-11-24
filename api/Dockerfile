# Start with the official Golang base image
FROM golang:1.18-alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the Go files into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose the API port
EXPOSE 8080

# Command to run the API
CMD ["./main"]