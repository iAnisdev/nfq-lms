
FROM golang:1.20-alpine3.17

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum .env ./
RUN go mod download

COPY * ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /lms

EXPOSE 4000

# Run
CMD ["/lms"]