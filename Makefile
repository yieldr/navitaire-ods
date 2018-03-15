VERSION ?= $(shell git describe --tags)

PKG = github.com/yieldr/navitaire-ods
PKGS = $(shell go list ./... | grep -v /vendor/ | grep -v /test)

LDFLAGS = "-s -w -X github.com/yieldr/navitaire-ods/pkg/version.Version=$(VERSION)"

OS ?= darwin
ARCH ?= amd64

build:
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/navitaire-ods-$(OS)-$(ARCH) -ldflags $(LDFLAGS)

test:
	@go test $(PKGS)

vet:
	@go vet $(PKGS)

generate:
	@go generate $(PKGS)
