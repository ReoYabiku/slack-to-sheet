build:
	docker compose build --no-cache
up:
	docker compose up -d
	docker compose exec golang /bin/sh -c "go run ./cmd"
down:
	docker compose down
sh:
	docker compose exec golang /bin/sh