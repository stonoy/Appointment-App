-- +goose Up
create table patient(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    age int not null,
    gender text not null,
    user_id uuid not null unique
    references users(id)
    on delete cascade
);

-- +goose Down
drop table patient;