package db

import (
	"context"
	test "github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    "Jaine",
		Balance:  1000,
		Currency: "REAL",
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	test.NoError(t, err)
	test.NotEmpty(t, account)

	test.Equal(t, args.Owner, account.Owner)
	test.Equal(t, args.Balance, account.Balance)
	test.Equal(t, args.Currency, account.Currency)
}
