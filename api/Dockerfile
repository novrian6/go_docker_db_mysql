# Use a newer Golang version that supports Go 1.23
FROM golang:1.23-alpine

# Set the current working directory inside the container
WORKDIR /usr/src/app

# Copy the Go module and sum files
COPY go.mod go.sum ./

# Download Go module dependencies (this will be cached if the go.mod and go.sum files are not changed)
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the API will run on
EXPOSE 8080

# Run the Go app
CMD ["./main"]
