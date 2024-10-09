swagger:
	@swag init

mock:
	@go generate ./...

test:
	@go test -v ./...

build:
	@go build -o bin/main main.go

run:
	@go run main.go
