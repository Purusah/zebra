PROJECTNAME=$(shell basename "$(PWD)")
GOVERSION=1.15
GOLINTVERSION = v1.31.0

.PHONY: build test

build:
	docker build -f $(PWD)/build/docker/Dockerfile --tag $(PROJECTNAME):latest --target dev .

lint:
	docker run -t --rm -v $(shell pwd):/$(PROJECTNAME) -w /$(PROJECTNAME) golangci/golangci-lint:$(GOLINTVERSION) golangci-lint run -v

fmt: build
	docker run -it --rm -v $(shell pwd):/$(PROJECTNAME) $(PROJECTNAME):latest go fmt -x -n

test: build
	docker run -it --rm $(PROJECTNAME):latest go test -v -race -tags=all -coverprofile cp.out ./...

test-unit: build
	docker run -it --rm $(PROJECTNAME):latest go test -v -race -tags=unit -coverprofile cp.out ./...

test-integration: build
	docker run -it --rm $(PROJECTNAME):latest go test -v -race -tags=integration ./...
