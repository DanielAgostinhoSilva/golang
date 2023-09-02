-- +goose Up
CREATE TABLE partner
(
    id         CHAR(36) PRIMARY KEY,
    name       VARCHAR(255)
);

-- +goose Down
DROP TABLE partner;
