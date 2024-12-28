FROM golang:1.23.4-alpine

RUN apk update && \
    apk upgrade

WORKDIR /app

COPY . /app 

RUN go mod tidy

RUN go install golang.org/x/tools/cmd/goimports@latest

