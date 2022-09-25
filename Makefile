DOCKER_REGISTRY=socheatsok78
DOCKER_NAME=libreoffice-unoserver-rest-api
DOCKER_TAG=nightly
DOCKER_IMAGE=${DOCKER_REGISTRY}/${DOCKER_NAME}:${DOCKER_TAG}

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

docker-build: build-linux
	docker build --pull --rm -f "Dockerfile" -t ${DOCKER_IMAGE} "."

docker-run:
	docker run -it --rm  -p "2004:2003" \
		${DOCKER_IMAGE}

clean:
	@rm -rf bin
