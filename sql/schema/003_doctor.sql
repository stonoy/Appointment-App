-- +goose Up
create table doctor(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    specialty text not null,
    license_number text not null unique,
    user_id uuid not null unique
    references users(id)
    on delete cascade
);

-- +goose Down
drop table doctor;