-- +goose Up
create table availability(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    location text not null,
    timing timestamp not null,
    duration int not null,
    max_patient int not null,
    current_patient int not null,
    treatment text not null,
    doctor_id uuid not null
    references doctor(id)
    on delete cascade
);

-- +goose Down
drop table availability;