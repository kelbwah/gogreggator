run:
	@./bin/gogreggator

build:
	@go build -o ./bin/gogreggator ./cmd/app/main.go

test:
	@go test -v ./...
