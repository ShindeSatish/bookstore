# Start from a lightweight Golang image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's code
COPY . .

# Copy Swagger docs (assuming docs are in the root directory of your project)
COPY docs/ /app/docs/

# Build the application
RUN go build -o main ./main.go

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]
