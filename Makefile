#!/usr/bin/env make -f

export GOPATH := $(CURDIR):$(CURDIR)/vendor
export PATH := $(CURDIR)/bin:$(PATH)

.PHONY: bin/aggremator bin/aggremator.linux

default: bin/aggremator

deploy: bin/aggremator.linux
	scp bin/aggremator.linux camlistore:aggremator

bin/aggremator.linux:
	GOOS=linux go build -o bin/aggremator.linux aggremator

bin/aggremator:
	GOOS=darwin go build -o bin/aggremator aggremator

doc:
	godoc -http :6060
