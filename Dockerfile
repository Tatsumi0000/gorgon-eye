FROM golang:1.23.4-alpine

RUN apk update && \
    apk upgrade

WORKDIR /app

COPY . /app 

RUN go install golang.org/x/tools/cmd/goimports@v0.28.0 && \ 
    go install github.com/spf13/cobra-cli@v1.3.0

RUN go mod tidy

