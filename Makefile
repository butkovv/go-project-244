build:
	go build -o bin/gendiff ./cmd/gendiff

lint:
	golangci-lint run ./...

test:
	go mod tidy
	go test -v ./...
