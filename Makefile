# Go config
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Config
TADATAKA_DIR = $(HOME)/.tadataka
BINARY_PATH=$(TADATAKA_DIR)/bin
BINARY_NAME=tadataka

build:
	mkdir -p $(TADATAKA_DIR)/bin
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v