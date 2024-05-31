-- name: CreateUser :one
insert into users(id, created_at, updated_at, name, email, password, role)
values ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: IsSetForAdmin :one
select
    case
    when count(*) = 0 then true
    else false
    end user_count_admin
from users;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetUserById :one
select * from users where id = $1;
