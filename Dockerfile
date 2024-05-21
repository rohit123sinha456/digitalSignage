# Stage 1: Build stage
FROM golang:1.22.3-alpine3.19 AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum .env ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o myapp .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .

# Set the timezone and install CA certificates
# RUN apk --no-cache add ca-certificates tzdata
EXPOSE 8080
# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]