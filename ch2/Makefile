.DEFAULT_GOAL := clean

.PHONY=format vet run clean
format:
	clear
	go fmt ./...

vet: format
	go vet ./...

run: vet
	go build -o build ./data_types.go && ./build

clean: run
	@echo
	rm -rf ./build