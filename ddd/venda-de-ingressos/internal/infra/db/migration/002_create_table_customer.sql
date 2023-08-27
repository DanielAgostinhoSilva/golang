-- +goose Up
CREATE TABLE customer
(
    id   CHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    cpf VARCHAR(255)
);

-- +goose Down
DROP TABLE customer;
