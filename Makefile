VERSION := $(shell git describe --long --tags)
TAG := $(shell git tag)
INSTALL := go install -ldflags "-X main.version=$(VERSION)" ./...
BINARY := envcrypt-$(shell uname -s)-$(shell uname -m)
CHECKSUM := $(BINARY).sha256
TARGETS := build install test

$(TARGETS):
	go $@

envcrypt:
	GOBIN=$(CURDIR) $(INSTALL)

binary: $(BINARY)
$(BINARY): envcrypt
	mv $< $@

checksum: $(CHECKSUM)
$(CHECKSUM): $(BINARY)
	shasum -p -a 256 $< > $@

release: $(BINARY) $(CHECKSUM)
	hub release create -a $(BINARY) -m $(TAG) -a $(CHECKSUM) $(TAG)

