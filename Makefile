run:
	go run cmd/main.go

db: postgres-image postgres-run wait-for-postgres migrate-up

# start the docker container which contains the postgresql image
start-db:
	-docker start financial_transaction_container

# gracefully shut down docker container which allows data to persists in postgresql
stop-db:
	-docker stop financial_transaction_container

# delete and create a new container - resets the database to default (for testing purposes)
restart-db: stop-db remove-db remove-volume db

# pulls postgres image
postgres-image:
	docker pull postgres:16-alpine

# runs postgres within docker container, allows connection via port :5432
postgres-run:
	docker run \
	--name financial_transaction_container \
	-p 5432:5432 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=secret \
	-e POSTGRES_DB=financial-transaction-db \
	-v pgdata:/var/lib/postgresql/data \
	-d postgres:16-alpine \

# waits for postgres container to be ready - without this, migrate will run immediately and fails
wait-for-postgres:
	@echo "Waiting for PostgreSQL to start..."
	@sleep 3;
	@until docker exec financial_transaction_container pg_isready -U root -h localhost -p 5432; do \
		sleep 5; \
	done
	@echo "PostgreSQL is ready"

# uses goose to migrate data into postgres database
migrate-up: wait-for-postgres
	goose -dir db postgres "postgresql://root:secret@localhost:5432/financial-transaction-db?sslmode=disable" up;

# deletes existing container
remove-db:
	-docker rm financial_transaction_container

# removes the volume pgdata that is mounted on host machine
remove-volume:
	-docker volume rm pgdata


# to generate coverage report and display in html
coverage:
	go test -coverprofile=coverage.out ./app/usecase
	go tool cover -html=coverage.out

# unit-tests for service layer
unit-test:
	go test -coverprofile=coverage.out ./app/usecase
	go test -coverprofile=coverage.out ./app/adapter/http/handlers

CURRENT_DIR := $(shell pwd)
swagger: swagger-delete
	docker run --name new-swagger-ui-container -p 80:8080 -e SWAGGER_JSON=/api.yaml -v $(CURRENT_DIR)/swagger.yaml:/api.yaml -d swaggerapi/swagger-ui

swagger-delete: swagger-stop-rm

swagger-stop-rm:
	-docker stop new-swagger-ui-container
	-docker rm new-swagger-ui-container

swagger-editor:
	docker run -p 8080:8080 swaggerapi/swagger-editor
