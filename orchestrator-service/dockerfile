# Use the official Golang image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Install air for live reloading
RUN go install github.com/cosmtrek/air@v1.51.0

# Copy Go module files to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose the port the Go app runs on
EXPOSE 8001

# The command to start the app with air for hot reloading 🔄
CMD ["air", "-c", ".air.toml"]