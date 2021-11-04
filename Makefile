ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif


run: fmt
	go run commandline/console.go

test: fmt
	go test -v -race

fmt:
	go fmt *.go
	go fmt commandline/*.go

modinit:
	go mod init github.com/siongui/goshorturl

modtidy:
	go mod tidy
