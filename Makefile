.PHONY: run build test clean migrate deps

run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

migrate:
	mysql -u root -p < docs/database.sql

deps:
	go mod download
	go mod tidy
