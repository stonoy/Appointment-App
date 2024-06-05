-- name: CreatePatient :one
insert into patient(id, created_at, updated_at, name, age, gender, user_id)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetPatientFromUserId :one
select * from patient where user_id = $1;

-- name: CheckUserIsPatient :one
select *
from patient where user_id = $1;