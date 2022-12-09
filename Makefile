POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DATABASE=medium_user_service_db

CURRENT_DIR=$(shell pwd)

-include .env
  
DB_URL="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"


print:
	echo "$(DB_URL)"
	
swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

local-up:
	docker compose --env-file ./.env.docker up -d

proto-gen:
	rm -rf genproto
	./scripts/gen-proto.sh ${CURRENT_DIR}
	
pull-sub-module:
	git submodule update --init --recursive

update-sub-module:
	git submodule update --remote --merge

.PHONY: start migrateup migratedown