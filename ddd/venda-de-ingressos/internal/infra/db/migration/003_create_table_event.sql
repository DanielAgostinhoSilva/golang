-- +goose Up
CREATE TABLE event
(
    id                   CHAR(36) PRIMARY KEY,
    name                 VARCHAR(255),
    description          VARCHAR(255),
    date                 TIMESTAMP,
    published            TINYINT(1),
    total_spots          INTEGER,
    total_spots_reserved INTEGER,
    partner_id           CHAR(36)
);

ALTER TABLE event
    ADD CONSTRAINT fk_partner_event
        foreign key (partner_id) references partner (id);

-- +goose Down
DROP TABLE event;
