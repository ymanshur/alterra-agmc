# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine AS builder
LABEL maintainer="Yusuf Manshur<ymanshur@gmail.com>"

WORKDIR /app

# Copy all app files
COPY . ./

# Download necessary Go modules
# COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# COPY *.go ./

# Build executable binary
RUN go build -o /main

# Build a small image
FROM alpine:3.16.0

WORKDIR /

COPY --from=builder /main .

EXPOSE 8080

ENTRYPOINT ["./main"]
