run:
	go run main.go

debug:
	go run main.go -v

build:
	go build

test:
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

test-debug:
	go test -cover -v -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html
