-- +goose Up
-- +goose StatementBegin

CREATE TABLE "account"
(
    id INTEGER NOT NULL PRIMARY KEY,
    balance DECIMAL(20,10) NOT NULL
);

CREATE TABLE transaction
(
    transaction_id SERIAL PRIMARY KEY,
    source_account_id INTEGER NOT NULL,
    destination_account_id INTEGER NOT NULL,
    amount DECIMAL(20, 10)  NOT NULL
);

INSERT INTO account (id, balance)
VALUES (123, 100.123456789),
       (124, 200.23344),
       (125, 400.23344);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
DROP TABLE transaction;
-- +goose StatementEnd
