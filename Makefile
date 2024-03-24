build:
	@go build -o bin/basic-ecom cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go