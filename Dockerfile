# Start from Golang base image
FROM golang:latest

# Metadata (example)
LABEL version="1.0"
LABEL description="This is a Go web application for ASCII art."
LABEL maintainer="baratovarslon8@gmail.com"
LABEL author="mnozimjo"
LABEL project="ascii-art-web-dockerize"

# Set the working directory inside the container
WORKDIR /app


# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o ascii-art-web-dockerize .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
CMD ["./ascii-art-web-dockerize"]
