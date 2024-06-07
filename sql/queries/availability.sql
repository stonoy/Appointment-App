-- name: CreateAvailability :one
insert into availability(id, created_at, updated_at, location, timing, duration, max_patient, current_patient, treatment, doctor_id)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning *;

-- name: UpdateCurrentPatient :one
update availability
set updated_at = NOW(),
current_patient = current_patient + 1
where id = $1 and timing > NOW() and current_patient < max_patient
returning *;

-- name: DelCurrentPatient :one
update availability
set updated_at = NOW(),
current_patient = current_patient - 1
where id = $1 and timing > NOW() and current_patient > 0
returning *;

-- name: GetAvailability :many
select a.*,d.name,d.specialty from availability a
inner join doctor d
on a.doctor_id = d.id
where location like $1 and timing between $2 and $3
limit $4
offset $5;

-- name: GetNumAvailability :one
select count(*) from availability
where location like $1 and timing between $2 and $3;

-- name: GetAvailabilityDoctor :many
select * from availability
where location like $1 and timing between $2 and $3 and doctor_id = $4
limit $5
offset $6;

-- name: GetNumAvailabilityDoctor :one
select count(*) from availability
where location like $1 and timing between $2 and $3 and doctor_id = $4;