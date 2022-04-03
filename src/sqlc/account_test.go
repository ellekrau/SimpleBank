package db

import (
	"context"
	"database/sql"
	"github.com/ellekrau/SimpleBank/utils/random"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Accounts {
	args := CreateAccountParams{
		Owner:    random.Owner(),
		Balance:  random.Amount(),
		Currency: random.Currency(),
	}
	resultAccount, err := db.CreateAccount(ctx, args)

	require.NoError(t, err)
	require.NotEmpty(t, resultAccount)

	require.NotZero(t, resultAccount.ID)
	require.NotZero(t, resultAccount.CreatedAt)

	require.Equal(t, args.Balance, resultAccount.Balance)
	require.Equal(t, args.Owner, resultAccount.Owner)
	require.Equal(t, args.Currency, resultAccount.Currency)

	return resultAccount
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	dbAccount := createRandomAccount(t)
	resultAccount, err := db.GetAccount(ctx, dbAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resultAccount)

	require.Equal(t, dbAccount, resultAccount)
}

func TestUpdateAccount(t *testing.T) {
	dbAccount := createRandomAccount(t)
	updateAccountParams := UpdateAccountParams{
		Balance: random.Amount(),
		ID:      dbAccount.ID,
	}
	resultAccount, err := db.UpdateAccount(ctx, updateAccountParams)

	require.NoError(t, err)
	require.NotEmpty(t, resultAccount)

	require.Equal(t, dbAccount.ID, resultAccount.ID)
	require.Equal(t, dbAccount.Owner, resultAccount.Owner)
	require.Equal(t, dbAccount.Currency, resultAccount.Currency)
	require.WithinDuration(t, dbAccount.CreatedAt, resultAccount.CreatedAt, time.Second)

	require.Equal(t, updateAccountParams.Balance, resultAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	dbAccount := createRandomAccount(t)
	err := db.DeleteAccount(ctx, dbAccount.ID)
	require.NoError(t, err)

	resultAccount, err := db.GetAccount(context.Background(), dbAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, resultAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	listAccountsParams := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	resultAccounts, err := db.ListAccounts(ctx, listAccountsParams)

	require.NoError(t, err)
	require.Len(t, resultAccounts, 5)

	for _, account := range resultAccounts {
		require.NotEmpty(t, account)
	}
}
