run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

docker-up:
	docker-compose up -d

migrate:
	echo "Run migrations"

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down
