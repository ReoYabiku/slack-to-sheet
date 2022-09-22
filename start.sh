#!/bin/sh
docker compose up -d
docker compose exec golang /bin/sh -c "go run ./cmd"
docker compose exec golang /bin/sh