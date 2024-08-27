build:
	@go build -o bin/ecomm cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecomm