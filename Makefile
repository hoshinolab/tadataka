# Go config
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Config
BINARY_PATH=./bin
BINARY_NAME=tadataka

build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v