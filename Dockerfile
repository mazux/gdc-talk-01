FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./binary ./cmd/app/main.go

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates
WORKDIR /tmp

COPY --from=builder /tmp/app/binary ./binary

# Run the binary program produced by `go install`
CMD ["/binary"]
