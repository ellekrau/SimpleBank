package db

import (
	"database/sql"
	"github.com/ellekrau/SimpleBank/utils/random"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createNewTransferAccountID(t *testing.T) sql.NullInt64 {
	return sql.NullInt64{
		Int64: createRandomAccount(t).ID,
		Valid: true,
	}
}

func createRandomTransfer(t *testing.T) Transfers {
	fromAccountID := createNewTransferAccountID(t)
	toAccountID := createNewTransferAccountID(t)
	amount := random.Amount()

	createTransferParams := CreateTransferParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	resultTransfer, err := db.CreateTransfer(ctx, createTransferParams)

	require.NoError(t, err)
	require.NotEmpty(t, resultTransfer)

	require.NotZero(t, resultTransfer.ID)
	require.Equal(t, fromAccountID, resultTransfer.FromAccountID)
	require.Equal(t, toAccountID, resultTransfer.ToAccountID)
	require.Equal(t, amount, resultTransfer.Amount)
	require.NotZero(t, resultTransfer.CreatedAt)

	return resultTransfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	dbTransfer := createRandomTransfer(t)
	resultTransfer, err := db.GetTransfer(ctx, dbTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resultTransfer)

	require.Equal(t, dbTransfer, resultTransfer)
}

func TestUpdateTransfer(t *testing.T) {
	dbTransfer := createRandomTransfer(t)
	updateTransferParams := UpdateTransferParams{
		FromAccountID: dbTransfer.FromAccountID,
		ToAccountID:   createNewTransferAccountID(t),
		Amount:        random.Amount(),
		ID:            dbTransfer.ID,
	}
	resultTransfer, err := db.UpdateTransfer(ctx, updateTransferParams)

	require.NoError(t, err)
	require.NotEmpty(t, resultTransfer)

	require.Equal(t, dbTransfer.FromAccountID, resultTransfer.FromAccountID)
	require.WithinDuration(t, dbTransfer.CreatedAt, resultTransfer.CreatedAt, time.Second)

	require.Equal(t, updateTransferParams.ToAccountID, resultTransfer.ToAccountID)
	require.Equal(t, updateTransferParams.Amount, resultTransfer.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	dbTransfer := createRandomTransfer(t)
	err := db.DeleteTransfer(ctx, dbTransfer.ID)
	require.NoError(t, err)

	resultTransfer, err := db.GetTransfer(ctx, dbTransfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, resultTransfer)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	listTransfersParams := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	resultTransfers, err := db.ListTransfers(ctx, listTransfersParams)

	require.NoError(t, err)
	require.Len(t, resultTransfers, 5)

	for _, transfer := range resultTransfers {
		require.NotEmpty(t, transfer)
	}
}
