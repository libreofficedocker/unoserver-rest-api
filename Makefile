VERSION=local
DOCKER_REGISTRY=libreoffice-docker
DOCKER_NAME=libreoffice-unoserver-rest-api
DOCKER_TAG=nightly
DOCKER_IMAGE=${DOCKER_REGISTRY}/${DOCKER_NAME}:${DOCKER_TAG}

OUTPUT := build

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	SHA_CMD := shasum -a 256
else
	SHA_CMD := sha256sum
endif

install:
	@go mod tidy

run:
	@go run unoserver-rest-api.go

build:
	$(call go-build,linux,amd64)
	$(call go-build,linux,arm64)

build-darwin:
	$(call go-build,darwin,amd64)
	$(call go-build,darwin,arm64)

build-docker: build
	DOCKER_BUILDKIT=1 docker build --rm -f "Dockerfile" -t ${DOCKER_IMAGE} "."

run-docker:
	docker run -it --rm  -p "2003:2003" ${DOCKER_IMAGE}

clean:
	@rm -rf $(OUTPUT); true

# define function
define go-build
	@echo "- Building for $(1)-$(2)..."
	@echo
	@CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -ldflags="-s -w -X main.Version=${VERSION}" -o $(OUTPUT)/unoserver-rest-api-$(1)-$(2) unoserver-rest-api.go
	@upx $(OUTPUT)/unoserver-rest-api-$(1)-$(2)
	@cd $(OUTPUT) && $(SHA_CMD) unoserver-rest-api-$(1)-$(2) > unoserver-rest-api-$(1)-$(2).sha256
endef
