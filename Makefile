install:
	@go mod tidy

run:
	@go run server.go

build: build-linux build-darwin

build-linux:
	GOOS=linux go build -ldflags="-s -w" -o bin/server-linux server.go
	upx bin/server-linux

build-darwin:
	GOOS=darwin go build -ldflags="-s -w" -o bin/server-darwin server.go
	upx bin/server-darwin
