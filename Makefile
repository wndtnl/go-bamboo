GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_FMT=$(GO_CMD) fmt
GO_MOD=$(GO_CMD) mod
GO_CLEAN=$(GO_CMD) clean

all: clean deps fmt test build

build:
	$(GO_BUILD) ./...

test:
	$(GO_TEST) ./... -v

fmt:
	$(GO_FMT) ./...

deps:
	$(GO_MOD) download

clean:
	$(GO_CLEAN) ./...
