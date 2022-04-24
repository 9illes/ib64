BINARY_DIR=bin/
BINARY_NAME=ib64

.PHONY: all

all: clean build build_slim xbuild

build:
	go build -o $(BINARY_DIR)$(BINARY_NAME) main.go
	file $(BINARY_DIR)$(BINARY_NAME)

build_slim:
	go build -ldflags "-s -w" -o $(BINARY_DIR)$(BINARY_NAME)_slim main.go
	file $(BINARY_DIR)$(BINARY_NAME)_slim

xbuild:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_DIR)$(BINARY_NAME).exe main.go
	GOOS=windows GOARCH=386 go build -o $(BINARY_DIR)$(BINARY_NAME)_dwn main.go

test:
	go test -v ./base64

clean:
	go clean
	rm -f $(BINARY_DIR)$(BINARY_NAME) $(BINARY_DIR)$(BINARY_NAME).exe $(BINARY_DIR)$(BINARY_NAME)_dwn $(BINARY_DIR)$(BINARY_NAME)_slim
