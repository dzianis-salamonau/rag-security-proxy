BINARY := bin/server

.PHONY: build run test lint docker-build compose-up compose-down clean

build:
	CGO_ENABLED=0 go build -o $(BINARY) ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

lint:
	go vet ./...

docker-build:
	docker build -t rag-security-proxy:dev .

compose-up:
	docker compose up -d --build

compose-down:
	docker compose down -v

clean:
	rm -rf bin
