-- +goose Up
CREATE TABLE user
(
    id         bigserial PRIMARY KEY,
    name       character varying(200) NOT NULL,
    age        smallint               NOT NULL DEFAULT 0 CHECK (age >= 0 AND age <= 200),
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT into user(name, age)
VALUES ('Fong', 23);

INSERT
into user (name, age)
VALUES ('Kevin', 40);

INSERT
into user (name, age)
VALUES ('Mary', 18);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down

DROP TABLE IF EXISTS gateways;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
