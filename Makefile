VERSION=local
DOCKER_REGISTRY=socheatsok78
DOCKER_NAME=libreoffice-unoserver-rest-api
DOCKER_TAG=nightly
DOCKER_IMAGE=${DOCKER_REGISTRY}/${DOCKER_NAME}:${DOCKER_TAG}

OUTPUT := output
OUTPUT := $(abspath $(OUTPUT))

install:
	@go mod tidy

run:
	@go run unoserver-rest-api.go

build: build-linux build-darwin

build-linux:
	GOOS=linux go build -ldflags="-s -w -X main.Version=${VERSION}" -o bin/unoserver-rest-api-linux unoserver-rest-api.go
	upx bin/unoserver-rest-api-linux
	mkdir -p rootfs/usr/bin
	cp bin/unoserver-rest-api-linux rootfs/usr/bin/unoserver-rest-api

build-darwin:
	GOOS=darwin go build -ldflags="-s -w -X main.Version=${VERSION}" -o bin/unoserver-rest-api-darwin unoserver-rest-api.go
	upx bin/unoserver-rest-api-darwin

docker-build: build-linux
	docker build --pull --rm -f "Dockerfile" -t ${DOCKER_IMAGE} "."

docker-run:
	docker run -it --rm  -p "2003:2003" \
		${DOCKER_IMAGE}

s6-overlay-module: $(OUTPUT)/s6-overlay-module.tar.zx

$(OUTPUT)/s6-overlay-module.tar.zx:
	exec mkdir -p $(OUTPUT)
	cd rootfs && tar -Jcvf $@ --owner=0 --group=0 --numeric-owner .

clean:
	rm -rf bin; true
	rm -rf $(OUTPUT); true
