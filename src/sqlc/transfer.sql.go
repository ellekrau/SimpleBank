// Code generated by sqlc. DO NOT EDIT.
// source: transfer.sql

package db

import (
	"context"
	"database/sql"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID sql.NullInt64 `json:"fromAccountID"`
	ToAccountID   sql.NullInt64 `json:"toAccountID"`
	Amount        int64         `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE id = $1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers ORDER BY id LIMIT $1 OFFSET $2
`

type ListTransfersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfers, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfers
	for rows.Next() {
		var i Transfers
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers SET from_account_id = $1, to_account_id = $2, amount = $3  where id = $4 RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	FromAccountID sql.NullInt64 `json:"fromAccountID"`
	ToAccountID   sql.NullInt64 `json:"toAccountID"`
	Amount        int64         `json:"amount"`
	ID            int64         `json:"id"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
		arg.ID,
	)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
