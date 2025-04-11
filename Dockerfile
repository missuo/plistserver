FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /go/src/github.com/missuo/plistserver

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o plistserver .

# Use a small alpine image for the final container
FROM alpine:latest
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /go/src/github.com/missuo/plistserver/plistserver .
# Expose the service port
EXPOSE 3788
# Run the binary
CMD ["/app/plistserver"]