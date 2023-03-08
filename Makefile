run:
	go run cmd/main.go

db:
	docker-compose up -d

test:
	go test ./...