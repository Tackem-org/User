run:
	@go run main.go

build:
	@go build

test:
	@FILES="$(shell go list ./... | grep -v /static)"; go test -timeout 30s -cover -coverprofile=cover.out $$FILES
	@go tool cover -func=cover.out | tail -n 1

test-debug:
	@FILES="$(shell go list ./... | grep -v /static)"; go test -timeout 30s -cover -v $$FILES
	@go test -cover -v ./...

test-html:
	@FILES="$(shell go list ./... | grep -v /static)"; go test -timeout 30s -cover -coverprofile=cover.out $$FILES
	@go tool cover -html=cover.out -o coverage.html
