# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine

WORKDIR /app

# Download necessary Go modules
COPY ./ .
RUN go get
RUN apk update 
RUN apk add  gcc
RUN go get github.com/mattn/go-sqlite3
# RUN go mod init github.com/mattn/go-sqlite3
 CMD ["go","run","start.go"]

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...