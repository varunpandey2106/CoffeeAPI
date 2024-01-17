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
	docker run --name coffee_container -p 5433:5432 -e POSTGRES_USER=varunpandey -e POSTGRES_PASSWORD=boombamboom -d postgres:12-alpine

# Create the 'coffee_db' database in the PostgreSQL container
create_db:
	docker exec -it coffee_container createdb --username=varunpandey --owner=varunpandey coffee_db
