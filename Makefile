OS ?= darwin
ARCH ?= amd64

build: generate
	GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/navitaire-ods-$(OS)-$(ARCH)

generate:
	go generate ./...
