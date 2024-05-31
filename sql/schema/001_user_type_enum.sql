-- +goose Up
create type user_role as enum ('patient', 'doctor', 'admin');

-- +goose down
drop type user_role;