install:
	@go mod tidy

run:
	@go run unoserver-rest-api.go

build: clean build-linux build-darwin

build-linux:
	GOOS=linux go build -ldflags="-s -w" -o bin/unoserver-rest-api-linux unoserver-rest-api.go
	upx bin/unoserver-rest-api-linux

build-darwin:
	GOOS=darwin go build -ldflags="-s -w" -o bin/unoserver-rest-api-darwin unoserver-rest-api.go
	upx bin/unoserver-rest-api-darwin

clean:
	@rm -rf bin
