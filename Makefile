VERSION=local
DOCKER_REGISTRY=libreofficedocker
DOCKER_NAME=libreoffice-unoserver-rest-api
DOCKER_TAG=nightly
DOCKER_IMAGE=${DOCKER_REGISTRY}/${DOCKER_NAME}:${DOCKER_TAG}

OUTPUT := output
OUTPUT := $(abspath $(OUTPUT))

install:
	@go mod tidy

run:
	@go run unoserver-rest-api.go

build: build-bin build-rootfs
	
build-bin:
	$(call build,linux)
	$(call build,darwin)

build-rootfs:
	mkdir -p rootfs/usr/bin
	cp bin/unoserver-rest-api-linux rootfs/usr/bin/unoserver-rest-api

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

# define a reusable recipe
define build
	@echo "Building for $(1)..."
	CGO_ENABLED=0 GOOS=$(1) \
		go build -ldflags="-s -w -X main.Version=${VERSION}" -o bin/unoserver-rest-api-$(1) cli/unoserver-rest-api.go
		upx -k bin/unoserver-rest-api-$(1)
endef
