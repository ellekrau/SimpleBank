package db

import (
	"database/sql"
	"github.com/ellekrau/SimpleBank/utils/random"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createNewEntryAccountID(t *testing.T) sql.NullInt64 {
	return sql.NullInt64{
		Int64: createRandomAccount(t).ID,
		Valid: true,
	}
}

func createRandomEntry(t *testing.T) Entries {
	createEntryParams := CreateEntryParams{
		AccountID: createNewTransferAccountID(t),
		Amount:    random.Amount(),
	}
	resultEntry, err := db.CreateEntry(ctx, createEntryParams)

	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.NotZero(t, resultEntry.ID)
	require.Equal(t, createEntryParams.AccountID, resultEntry.AccountID)
	require.Equal(t, createEntryParams.Amount, resultEntry.Amount)
	require.NotZero(t, resultEntry.CreatedAt)

	return resultEntry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	dbEntry := createRandomEntry(t)
	resultEntry, err := db.GetEntry(ctx, dbEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.Equal(t, dbEntry, resultEntry)
}

func TestUpdateEntry(t *testing.T) {
	dbEntry := createRandomEntry(t)
	updateEntryParams := UpdateEntryParams{
		AccountID: createNewEntryAccountID(t),
		Amount:    random.Amount(),
		ID:        dbEntry.ID,
	}
	resultEntry, err := db.UpdateEntry(ctx, updateEntryParams)

	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.Equal(t, dbEntry.ID, resultEntry.ID)
	require.WithinDuration(t, dbEntry.CreatedAt, resultEntry.CreatedAt, time.Second)

	require.Equal(t, updateEntryParams.AccountID, resultEntry.AccountID)
	require.Equal(t, updateEntryParams.Amount, resultEntry.Amount)
}

func TestDeleteEntry(t *testing.T) {
	dbEntry := createRandomEntry(t)
	err := db.DeleteEntry(ctx, dbEntry.ID)
	require.NoError(t, err)

	resultEntry, err := db.GetEntry(ctx, dbEntry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, resultEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	listEntriesParams := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	resultEntries, err := db.ListEntries(ctx, listEntriesParams)

	require.NoError(t, err)
	require.Len(t, resultEntries, 5)

	for _, entry := range resultEntries {
		require.NotEmpty(t, entry)
	}
}
