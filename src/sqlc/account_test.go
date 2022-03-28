package db

import (
	"context"
	"github.com/ellekrau/SimpleBank/utils/enum/currency"
	test "github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    "Jaine",
		Balance:  1000,
		Currency: currency.Real,
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	test.NoError(t, err)
	test.NotEmpty(t, account)

	test.Equal(t, args.Owner, account.Owner)
	test.Equal(t, args.Balance, account.Balance)
	test.Equal(t, args.Currency, account.Currency)

	test.NotZero(t, account.ID)
	test.NotZero(t, account.CreatedAt)
}
