#BINARY ?= $(shell basename "$(PWD)")# binary name
BINARY := $(subst cmd/,,$(wildcard cmd/*))
CMD := $(wildcard cmd/*/*.go)
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

# Clean the build directory (before committing code, for example)
.PHONY: clean
clean: 
	rm -rv bin

PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

release: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o 'bin/$(os)-$(arch)/$(BINARY)' $(CMD)

.PHONY: release $(PLATFORMS)

