-- +migrate Up
CREATE TABLE users (id int);
-- +migrate Down
DROP TABLE users;
