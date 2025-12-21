TARGET = delegator
CGO_ENABLED = 0
GOOS = linux
GOARCH = amd64
GOAMD64 = v3
GO = CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOAMD64=$(GOAMD64) go
GOFLAGS = -ldflags "-s -w" -trimpath
GO_SOURCES = $(shell find . -type f -name '*.go')
WEB_SOURCES = $(shell find ./web/src ./web/public -type f)

.PHONY: all clean distclean

all: $(TARGET)

go.sum: $(GO_SOURCES)
	$(GO) mod tidy
	$(GO) fmt ./...
	touch $@

web/node_modules: ./web/package.json
	cd web && pnpm install
	touch $@

web/dist: web/node_modules $(WEB_SOURCES)
	cd web && pnpm run build
	touch $@

$(TARGET): go.sum web/dist
	$(GO) build -o $@ $(GOFLAGS) ./cmd/delegator

clean:
	rm -f $(TARGET)
	rm -rf web/dist

distclean: clean
	rm -rf web/node_modules
	rm -f delegator.db
