VERSION ?= $(shell git describe --tags)

PKG = github.com/yieldr/navitaire-ods
PKGS = $(shell go list ./... | grep -v /vendor/ | grep -v /test)

# SENTRY_DSN ?= "https://0ae4997acba549f18cf03b2ef7ce54d9@sentry.io/304990"

LDFLAGS = "-s -w -X github.com/yieldr/navitaire-ods/pkg/version.Version=$(VERSION)"

OS ?= darwin
ARCH ?= amd64

build:
ifeq ($(OS),windows)
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/navitaire-ods-$(OS)-$(ARCH).exe -ldflags $(LDFLAGS)
else
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/navitaire-ods-$(OS)-$(ARCH) -ldflags $(LDFLAGS)
endif

test:
	@go test $(PKGS)

vet:
	@go vet $(PKGS)

generate:
	@go generate $(PKGS)
