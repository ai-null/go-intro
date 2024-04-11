.DEFAULT_GOAL := clean

# reformat go
# check for any warnings & errors
# build and run
# clean build after run
.PHONY:fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o hello && ./hello

clean: build
	rm -rf hello