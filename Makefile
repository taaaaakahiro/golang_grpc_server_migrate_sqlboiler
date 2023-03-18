run:
	go run cmd/main.go

db:
	docker-compose up -d

test:
	go test ./... -v

sql:
	sqlboiler psql -c boiler.toml -o ./pkg/domain/models --wipe
