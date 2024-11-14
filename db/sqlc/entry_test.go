package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"githab.com/techschooll/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Accounts) Entries {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntriy(t *testing.T) {
	CreateRandomEntry(t, CreateRandomAccount(t))
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t, CreateRandomAccount(t))
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t, CreateRandomAccount(t))

	arg := UpdateEntryParams{
		ID:        entry1.ID,
		AccountID: entry1.AccountID,
		Amount:    util.RandomMoney(),
	}

	entry1, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry1.ID, entry1.ID)
	require.Equal(t, entry1.AccountID, entry1.AccountID)
	require.Equal(t, entry1.Amount, entry1.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t, CreateRandomAccount(t))

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, CreateRandomAccount(t))
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
