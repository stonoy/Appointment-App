-- +goose Up
create table appointment(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    status appointment_status not null,
    patient_id uuid not null
    references patient(id)
    on delete cascade,
    availability_id uuid not null
    references availability(id)
    on delete cascade,
    unique(patient_id, availability_id)
);

-- +goose Down
drop table appointment;