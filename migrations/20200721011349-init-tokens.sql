-- +migrate Up
alter table "users"
    add column "access_token" VARCHAR(255);
alter table "users"
    add column "refresh_token" VARCHAR(255);
-- +migrate Down
alter table "users"
    drop column "access_token";
alter table "users"
    drop column "refresh_token";
