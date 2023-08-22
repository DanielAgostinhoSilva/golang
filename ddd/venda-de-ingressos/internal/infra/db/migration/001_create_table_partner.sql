-- +goose Up
CREATE TABLE Partner
(
    Id   CHAR(36) PRIMARY KEY,
    Name VARCHAR(255)
);

-- +goose Down
DROP TABLE Partner;
