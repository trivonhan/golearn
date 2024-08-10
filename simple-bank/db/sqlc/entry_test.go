package db

import (
	"context"
	"database/sql"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams {
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t,arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryExpect := createRandomEntry(t, account)

	entryActual, err := testQueries.GetEntry(context.Background(), entryExpect.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryActual)

	require.Equal(t, entryExpect.ID, entryActual.ID)
	require.Equal(t, entryExpect.AccountID, entryActual.AccountID)
	require.Equal(t, entryExpect.Amount, entryActual.Amount)
	require.WithinDuration(t, entryExpect.CreatedAt, entryActual.CreatedAt, time.Second)
}

func TestGetListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		account := createRandomAccount(t)
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryOld := createRandomEntry(t, account)

	arg := UpdateEntryParams {
		ID: entryOld.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.ID, entry.ID)
	require.Equal(t, arg.Amount, entry.Amount)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryOld := createRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), entryOld.ID)

	entry, err := testQueries.GetEntry(context.Background(), entryOld.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, entry)
}