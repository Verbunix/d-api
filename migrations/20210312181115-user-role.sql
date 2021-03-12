-- +migrate Up
alter table users add column "role" VARCHAR(255);
create index "idx-users-role" on users (role);
-- +migrate Down
drop index "idx-users-role";
alter table users drop column "role";
