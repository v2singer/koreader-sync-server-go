APP_NAME = koreader-sync-server-go

.PHONY: build build-mac build-linux run

build: build-mac build-linux

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)-mac-arm64 main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux-amd64 main.go

run:
	go run main.go 