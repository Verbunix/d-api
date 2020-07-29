-- +migrate Up
alter table "users"
    add column "password" VARCHAR(255);
-- +migrate Down
alter table "users"
    drop column "password";
