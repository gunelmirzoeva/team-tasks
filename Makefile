
---
#### 7. **Add a simple `Makefile`**
It doesn’t have to do anything yet — placeholders are fine:
```makefile
run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

docker-up:
	docker-compose up -d

migrate:
	echo "Run migrations"
