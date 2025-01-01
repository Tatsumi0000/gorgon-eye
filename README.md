# 🐍 gorgon-eye [WIP]
RSS reader for CLI.

## 🧑‍💻 Develop
I develop with Docker, and we run Makefile to handle Docker commands with ease.

1. Create docker container.
```
make dev/setup
```
2. Run docker container and go program.
```
make dev/run
go run main.go
```

## 🧪 Test
```
make dev/run
go test -v ./...
```
