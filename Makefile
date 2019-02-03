#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_OUTPUT_FILE = confgen.exe
else
	BUILD_OUTPUT_FILE ?= confgen
endif

all: cleandep depend build

cleandep:
	@rm -rf vendor

depend:
	@dep ensure

build:
	@go build -o $(BUILD_OUTPUT_FILE) cmd/confgen/main.go

test:
	@echo "Running tests..."
	@go test `go list ./... | grep -v vendor/`
