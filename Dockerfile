# Base image
FROM golang:1.21-alpine

# Workdir
WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Source code
COPY . .

# Build
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run app
CMD ["./main"]
