BIN_NAME := shiori
PKG := github.com/AobaIwaki123/shiori/cmd/shiori
PREFIX ?= $(shell go env GOPATH)
BINDIR ?= $(PREFIX)/bin

.PHONY: all build install test verify clean

all: build

build:
	go build -o bin/$(BIN_NAME) ./cmd/shiori

install:
	go install ./cmd/shiori
	@echo "Installed $(BIN_NAME) to $(BINDIR)/$(BIN_NAME)"

test:
	go test -v ./...

verify:
	./scripts/verify-all.sh

clean:
	rm -rf bin/
