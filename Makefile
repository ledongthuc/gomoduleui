GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run 
GOTEST=$(GOCMD) test
BINARY_NAME=gomoduleui
DOCKER_IMG_NAME=$(BINARY_NAME)

all: test build
test:
	$(GOTEST) ./...
build: 
	$(GOBUILD) -o ./output/$(BINARY_NAME) -v ./main.go
run:
	$(GORUN) ./main.go
