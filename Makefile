.PHONY: run test build

run:
	go run cmd/main.go

tidy:
	go mod tidy

build-fyne:
	CGO_ENABLED=1 go build -o gestorvideo ./cmd/main.go