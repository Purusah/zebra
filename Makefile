PROJECTNAME=$(shell basename "$(PWD)")
GOVERSION=1.15
GOLINTVERSION = v1.31.0

.PHONY: build test

build:
	docker build -f $(PWD)/build/docker/Dockerfile --tag $(PROJECTNAME):latest --target dev .

lint:
	docker run -t --rm -v $(shell pwd):/$(PROJECTNAME) -w /$(PROJECTNAME) golangci/golangci-lint:$(GOLINTVERSION) golangci-lint run -v

test: build
	docker run -it --rm $(PROJECTNAME):latest go test -v -race -coverprofile cp.out ./...
