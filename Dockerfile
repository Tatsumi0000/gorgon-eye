FROM golang:1.23.4-alpine

RUN apk update && \
    apk upgrade

WORKDIR /app

COPY . /app 

RUN go install golang.org/x/tools/cmd/goimports && \ 
    go install github.com/spf13/cobra-cli

RUN go mod tidy

