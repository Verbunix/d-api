-- +migrate Up
create table users
(
    id         serial PRIMARY KEY,
    email      VARCHAR(255) UNIQUE NOT NULL,
    name       VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT current_timestamp
);

create index on users (email);
create index on users (name);
create index on users (created_at);

-- +migrate Down
DROP TABLE users;
