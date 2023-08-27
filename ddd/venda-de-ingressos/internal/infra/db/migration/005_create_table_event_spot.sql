-- +goose Up
CREATE TABLE event_spot
(
    id        CHAR(36) PRIMARY KEY,
    location  VARCHAR(255),
    reserved  TINYINT(1),
    published TINYINT(1),
    event_section_id CHAR(36)
);

ALTER TABLE event_spot
    ADD CONSTRAINT fk_event_spot_section
        foreign key (event_section_id) references event_section (id);

-- +goose Down
DROP TABLE event_spot;
