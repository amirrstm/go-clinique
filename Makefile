.PHONY: clean test security build run

APP_NAME = go-clinique
BUILD_DIR = ./build
MIGRATIONS_FOLDER = ./platform/migrations
DB_NAME = amir
DB_USER = amir
DB_PASS = 
DATABASE_URL = postgres://$(DB_USER):$(DB_PASS)@localhost/$(DB_NAME)?sslmode=disable

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	swag init

build: swag clean
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down 

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

seed:
	PGPASSWORD=$(DB_PASS) psql -h localhost -p 5432 -U$(DB_USER) -d $(DB_NAME) -a -f platform/seeds/001_seed_user_table.sql

docker.run: docker.setup docker.postgres docker.fiber migrate.up
	@echo "\n===========FGB==========="
	@echo "App is running...\nVisit: http://localhost:9000 OR http://localhost:9000/swagger/"

docker.setup:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network
	docker volume create go-clinique-pgdata

docker.fiber.build: swag
	docker build -t go-clinique:latest .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name go-clinique \
		--network dev-network \
		-p 9000:9000 \
		go-clinique

docker.postgres:
	docker run --rm -d \
		--name fibergb-postgres \
		--network dev-network \
		-e POSTGRES_USER=dev \
		-e POSTGRES_PASSWORD=dev \
		-e POSTGRES_DB=dev \
		-v fibergb-pgdata:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.stop: docker.stop.fiber docker.stop.postgres

docker.stop.fiber:
	docker stop fibergb-api || true

docker.stop.postgres:
	docker stop fibergb-postgres || true

docker.dev:
	docker-compose up