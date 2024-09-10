-- name: GetAdminById :one
SELECT * FROM Admin
WHERE id = sqlc.arg(admin_id) LIMIT 1;


-- name: DeleteAdminLoginById :exec
UPDATE Admin
SET isdelete = 1
WHERE id = sqlc.arg(admin_id);
