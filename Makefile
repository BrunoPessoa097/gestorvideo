.PHONY: run test build

run:
	go run cmd/main.go

tidy:
	go mod tidy