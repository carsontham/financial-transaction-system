-- +goose Up
-- +goose StatementBegin

CREATE TABLE "account"
(
    id SERIAL PRIMARY KEY,
    balance DECIMAL(15,8) NOT NULL
);

CREATE TABLE transaction
(
    transaction_id INTEGER PRIMARY KEY,
    source_account_id INTEGER NOT NULL,
    destination_account_id INTEGER NOT NULL,
    amount DECIMAL(15, 8)  NOT NULL
);

INSERT INTO account (id, balance)
VALUES (123, 100.23344),
       (124, 200.23344),
       (125, 400.23344);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
DROP TABLE transaction;
-- +goose StatementEnd
