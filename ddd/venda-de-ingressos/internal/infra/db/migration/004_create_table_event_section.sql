-- +goose Up
CREATE TABLE event_section
(
    id                  CHAR(36) PRIMARY KEY,
    name                VARCHAR(255),
    description         VARCHAR(255),
    published           TINYINT(1),
    total_spot          INTEGER,
    total_spot_reserved INTEGER,
    price               FLOAT,
    event_id            CHAR(36)
);

ALTER TABLE event_section
    ADD CONSTRAINT fk_event_section
        foreign key (event_id) references event (id);

-- +goose Down
DROP TABLE event_section;
