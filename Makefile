BIN_NAME := memo
PKG := github.com/AobaIwaki123/memo-cli/cmd/memo
PREFIX ?= $(shell go env GOPATH)
BINDIR ?= $(PREFIX)/bin

.PHONY: all build install test verify clean

all: build

build:
	go build -o bin/$(BIN_NAME) ./cmd/memo

install:
	go install ./cmd/memo
	@echo "Installed $(BIN_NAME) to $(BINDIR)/$(BIN_NAME)"

test:
	go test -v ./...

verify:
	./scripts/verify-all.sh

clean:
	rm -rf bin/
