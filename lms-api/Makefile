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
	@export $$(grep -v '^#' ../.env | xargs) && \
	mysql -h 127.0.0.1 -P 3310 -u root -p"$$MYSQL_ROOT_PASSWORD" "lms-backend" < seed.sql

ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}
