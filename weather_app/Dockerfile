# Use the official Go image as a parent image
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myweatherapp cmd/weather_app/main.go

# Use the scratch base image as the final base image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary from the builder image to the final image
COPY --from=builder /app/myweatherapp /

# Expose a port that your Go application will listen on
EXPOSE 8080

# Run the Go application when the container starts
CMD ["/myweatherapp"]