.PHONY: dev generate docker-up docker-down seed

dev:
	go run ./cmd/...

generate:
	go generate ./ent/...

docker-up:
	docker compose up -d

docker-down:
	docker compose down

seed:
	go run ./cmd/seed/...

ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}
