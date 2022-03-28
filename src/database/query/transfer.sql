-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1;

-- name: ListTransfers :many
SELECT * FROM transfers ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers SET from_account_id = $1 and to_account_id = $2 and amount = $3  where id = $4 RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;