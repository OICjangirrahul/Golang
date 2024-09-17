FROM golang:1.23

WORKDIR /usr/src/app

# Install air for live reloading
RUN go install github.com/air-verse/air@latest


# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -v -o /usr/local/bin/app ./...

# Command to run the application
CMD ["air", "-c", ".air.toml"]
