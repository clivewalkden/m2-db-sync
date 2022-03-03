.DEFAULT_GOAL := build

GO_BIN=${HOME}/go/go1.16.14/bin/go

fmt:
	${GO_BIN} fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	${GO_BIN} vet ./...
.PHONY:vet

build: vet
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=amd64 ${GO_BIN} build -o bin/main-freebsd-amd64 main.go
	GOOS=darwin GOARCH=arm64 ${GO_BIN} build -o bin/m2-db-sync-macos-arm64 main.go
	GOOS=linux GOARCH=amd64 ${GO_BIN} build -o bin/m2-db-sync-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 ${GO_BIN} build -o bin/m2-db-sync-windows-amd64.exe main.go
.PHONY:build