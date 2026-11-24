-- +goose Up
CREATE TABLE account (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL
);

CREATE TABLE role (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE CHECK (name IN ('admin', 'user', 'controller'))
);

INSERT INTO role(name) VALUES
    -- ('admin'),
    ('user'),
    ('controller');

CREATE TABLE account_role (
    account_id bigint NOT NULL REFERENCES account (id),
    role_id int NOT NULL REFERENCES role (id),
    PRIMARY KEY(account_id, role_id) 
);


-- +goose Down
DROP TABLE IF EXISTS account_role;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS account;
