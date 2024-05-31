-- +goose Up
create type user_role as enum ('patient', 'doctor', 'admin');

-- +goose Down
drop type user_role;