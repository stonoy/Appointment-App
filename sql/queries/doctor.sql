-- name: CreateDoctor :one
insert into doctor(id, created_at, updated_at, name, specialty,license_number, user_id)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetDoctorFromUserId :one
select * from doctor where user_id = $1;

-- name: GetAllDoctorsWithPage :many
select * from doctor
where specialty like $1 and name like $2
limit $3
offset $4;

-- name: GetNumOfAllDoctorsWithPage :one
select count(*) from doctor
where specialty like $1 and name like $2;