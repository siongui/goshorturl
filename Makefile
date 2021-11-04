ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif


test: fmt
	go test -v -race

fmt:
	go fmt *.go

modinit:
	go mod init github.com/siongui/goshorturl

modtidy:
	go mod tidy
