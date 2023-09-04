.PHONY: test run

test:
	gofmt -w . && go test -v ./pkg/... -cover -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

run:
	gofmt -w . && go run ./cmd/api/main.go