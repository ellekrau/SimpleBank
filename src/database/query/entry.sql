-- name: CreateEntry :one
INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries SET account_id = $1 and amount = $2 where id = $3 RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;