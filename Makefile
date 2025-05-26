build:
	@go build -o bin/lumbaumbah-backend cmd/main.go

run: build
	@./bin/lumbaumbah-backend

test:
	@go test -v ./...