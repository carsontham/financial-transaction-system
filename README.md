# Financial-Transaction-System
This application aims to demonstrate the following:
- Creating an application that adheres to [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- Dependency Injection to enable modular code
- Using Docker to spin up PostgresDB and Swagger UI
- Writing Unit Tests (>95% code coverage) through the use of mocking

# Introduction
The application follows [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and there are three main layers to this project:
- handler layer - interacts with all external interface (REST API)
- service/domain layer - interacts between handler and repository layer, conduct most business logics here
- repository layer - interacts with the database (where SQL queries can be found)

**Bootstrap**
- stores the logic for setting up server, api routes, custom validations

# To run the application: 
First ensure the database is set up. Run the command:
```
make set-db
```
This will create a postgreSQL DB using Docker. The data is migrated using Goose. Ensure you have Goose installed:

`go install github.com/pressly/goose/v3/cmd/goose@latest`

To ensure data persists within the container, run the following for graceful shutdown: 
```
make stop-db
``` 
To start the DB again, run :
```
make start-db
```
**Note: If you wish to "restart" the database with default data, you can still use `make set-db`** 

<h2>To run the server:</h2> 

```
make run
```
The server will listen on port 3000.

<h3>API Documentations </h3>
- To view the API documentations using OpenAPI, run:

```
make swagger
```
This runs swagger-ui using Docker on port :80. You can now visit localhost:80 to view the API documentations.

<img width="1000" alt="Screenshot 2024-04-20 at 5 49 38 PM" src="https://github.com/carsontham/financial-transaction-system/assets/127476216/77b42ca7-4fb4-4117-8977-1989ccd8c486">

<h3>API ENDPOINTS</h3>

**POST /accounts - create new account**

```
example req body :

        {
            "account_id": 123,
            "initial_balance": "100.23344"
        }

responses: 200, 400, 422, 500 
```

**GET /accounts/{account_id} - get an account by id**

```
example response body:

    {
        "status_code": 200,
        "data": {
            "account_id": 123,
            "balance": "100.123456789"
        }
    }

repsonse: 200, 404, 500
```

**POST /transactions - create a new payment transaction**

This request is made idempotent using an idempotency key. The assumption is that the Idempotency Key is generated on Client Side, for each new requests.

To test for Idempotency, include a unique string in request Header as
"X-Idempotency-Key". Subsequent requests with the same key will not perform a transfer but return the same results.
By default, if no key is present in header, a unique-key will be generated.

```
example req body:

        {
            "source_account_id": 123,
            "destination_account_id": 456,
            "amount": "100.12345"
        }

example response body:

        {
            "status_code": 200,
            "data": "transaction successful"
        }

repsonse: 200, 400, 404, 409, 422, 500

note - 409 occurs when balance is insufficient

```

**GET /transactions - Get all transactions**

This function is created mainly for easy retrieval of all the transactions without always going to the DB

```
example response body:

{
  "status_code": 200,
  "data": [
    {
      "transaction_id": 1,
      "source_account_id": 123,
      "destination_account_id": 321,
      "amount": "100.23344",
      "idempotency_key": "test-idempotency-key"
    }
  ]
}

repsonse: 200, 500
```


<h2> To access PostgreSQL DB on CLI: </h2>

```
docker exec -it {container-hash} psql -U root
```
CLI Commands in Postgres:
- Display Tables: `\l`
- Connect to DB: `\c dbname`
- List Tables: `\dt`
- Describe Table: `\d tablename`

<h2> Unit Tests </h2>

Unit tests have been added for handler and service layers. This was achieved using gomock. To run unit tests and check for its coverage:
```
make unit-test
```

