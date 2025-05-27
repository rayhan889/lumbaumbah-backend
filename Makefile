build:
	@go build -o bin/lumbaumbah-backend cmd/main.go

run: build
	@./bin/lumbaumbah-backend

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down