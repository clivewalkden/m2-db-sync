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
	GOOS=freebsd GOARCH=386 ${GO_BIN} build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 ${GO_BIN} build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 ${GO_BIN} build -o bin/main-windows-386.exe main.go
.PHONY:build