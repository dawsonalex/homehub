#BINARY ?= $(shell basename "$(PWD)")# binary name
SERVICE_NAME := $(subst cmd/,,$(wildcard cmd/*))
CMD := $(wildcard cmd/*/*.go)
PROTOC_FILE := $(wildcard cmd/*/schema/*.proto)
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

# Clean the build directory (before committing code, for example)
.PHONY: clean
clean: 
	rm -rv bin

PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

release: $(PLATFORMS)

build-protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        $(PROTOC_FILE)

$(PLATFORMS): build-protoc
	GOOS=$(os) GOARCH=$(arch) go build -o 'bin/$(os)-$(arch)/$(SERVICE_NAME)' $(CMD)

.PHONY: release $(PLATFORMS)

