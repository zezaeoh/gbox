
test: tidy
	go test -v ./...

build:
	go build -o gbox cmd/main.go

fmt: tidy
	go fmt ./...

tidy:
	go mod tidy