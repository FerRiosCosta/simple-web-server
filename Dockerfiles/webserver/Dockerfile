FROM golang:1.19.3-alpine

# Create working directory in /app
WORKDIR /app

# Copy all go config files.
COPY go.* ./

# Install required modules
RUN go mod download

# Copy source code
COPY *.go ./

# Run the Go build and locate the binary under /simple-web-server
RUN go build -o /simple-web-server

# Expose port to 8080
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/simple-web-server"]
