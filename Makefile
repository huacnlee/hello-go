run:
	@go run main.go monkey.go
test:
	@go test ./...
build:
	@go build -o build/hello-go main.go monkey.go