.DEFAULT_GOAL := build
fmt: 
	go fmt ./...
.PHONY:fmt
lint: fmt
	golint ./...
.PHONY:lint
vet: fmt
	go vet ./...
.PHONY:vet
	Makefiles |
lint-ci:
	golangci-lint run
.PHONY:lint-ci
build: vet
	mkdir build
	go build -o ./build/ ./cmd/main.go 
.PHONY:build
clean:
	rm -r ./build
.PHONY:clean