# Use the latest Go image
FROM golang:latest

# Set working directory inside the container
WORKDIR /app

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o /app/app main.go

# Set the correct CMD path
CMD ["/app/app"]
