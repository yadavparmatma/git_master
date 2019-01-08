# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=app
BINARY_DARWIN=$(BINARY_NAME)_darwin

all: test build

build:
	$(GOBUILD) -o ./out/$(BINARY_NAME)_darwin -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f ./out/$(BINARY_NAME)_darwin
run:
	$(GOBUILD) -o ./out/$(BINARY_NAME)_darwin -v
	./out/$(BINARY_NAME)_darwin

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./out/$(BINARY_DARWIN) -v
