# Build stage
FROM golang:1.22.5-alpine AS build

# Install Node.js and npm
RUN apk add --no-cache nodejs npm

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Install npm dependencies and build CSS
RUN npm install
RUN npm run build-css

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary and static files from the build stage
COPY --from=build /app/main .
COPY --from=build /app/static ./static
COPY --from=build /app/templates ./templates

RUN chmod +x /root/main

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
