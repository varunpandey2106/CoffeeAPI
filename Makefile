include .env

# Stop running Docker containers
stop_containers:
	@echo "Stopping other docker containers"
	@if [ $$(docker ps -q) ]; then \
		echo "Found and stopped other containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "No containers running..."; \
	fi

# Create a Docker container for PostgreSQL
create_container:
	docker run --name coffee_container -p 5432:5432 -e POSTGRES_USER=varunpandey -e POSTGRES_PASSWORD=boombamboom -d postgres:12-alpine

# Create the 'coffee_db' database in the PostgreSQL container
create_db:
	docker exec -it coffee_container createdb --username=varunpandey --owner=varunpandey coffee_db

#start container
start_container:
	docker start coffee_container

#make migrations
create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://varunpandey:boombamboom@localhost:5432/coffee_db?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://varunpandey:boombamboom@localhost:5432/coffee_db?sslmode=disable"


build:
	@if [ -f "coffee_api" ]; then \
		rm coffee_api; \
		echo "Deleted coffee_api"; \
	fi
	@echo "Building coffee_api..."
	go build -o coffee_api cmd/server/*.go

run: build
	@echo "Running coffee_api..."
	./coffee_api

stop:
	@echo "Stopping server..."
	@-pkill -SIGTERM -f "./coffee_api"
	@echo "Server stopped..."
