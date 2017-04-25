#!/usr/bin/env make -f

export GOPATH := $(CURDIR):$(CURDIR)/vendor
export PATH := $(CURDIR)/bin:$(PATH)

.PHONY: devel ./bin/server

default: devel

test: $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb test

bin/gb-vendor:
	go build -o bin/gb-vendor github.com/constabulary/gb/cmd/gb-vendor

bin/gb: bin/gb-vendor
	go build -o bin/gb github.com/constabulary/gb/cmd/gb

bin/% : $(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go) ./bin/gb
	gb build $( basename $@ )

doc:
	godoc -http :6060
