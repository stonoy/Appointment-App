-- name: CreateAppointment :one
insert into appointment(id, created_at, updated_at, status, patient_id, availability_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetAppointmentsPatient :many
select ap.*,av.location,av.timing,av.duration,av.treatment,d.name,d.specialty from appointment ap
inner join availability av
on ap.availability_id = av.id
inner join doctor d
on av.doctor_id = d.id
where status = $1 and location like $2 and treatment like $3 and timing between $4 and $5 and patient_id = $6
limit $7
offset $8;

-- name: GetNumAppointmentsPatient :one
select count(*) from appointment ap
inner join availability av
on ap.availability_id = av.id
inner join doctor d
on av.doctor_id = d.id
where status = $1 and location like $2 and treatment like $3 and timing between $4 and $5 and patient_id = $6;

-- name: DeleteAppointment :one
delete from appointment where patient_id = $1 and id = $2
returning *;