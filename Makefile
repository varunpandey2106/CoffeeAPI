include .env

stop_containers:
	@echo "Stopping other docker containers"
	@if [ $$(docker ps -q) ]; then \
		echo "Found and stopped other containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "No containers running..."; \
	fi

create_container:
		docker run --name coffee_container -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

