-- +goose Up
create type appointment_status as enum ('scheduled', 'completed', 'cancelled');

-- +goose Down
drop type appointment_status;