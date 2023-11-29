package db

import (
	"testing"
	"time"

	"github.com/CarlosVinicius3258/go-simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry{
	account := createRandomAccount(t)
	arg:= CreateEntryParams{
			AccountID: account.ID,
			Amount: util.RandomMoney(),
		}

	entry, err := testQueries.CreateEntry(GetContext(),arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
	
}

func TestGetEntry(t *testing.T){
	entry := createRandomEntry(t)
	output, err := testQueries.GetEntry(GetContext(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, entry.ID, output.ID)
	require.Equal(t, entry.AccountID, output.AccountID)
	require.Equal(t, entry.Amount, output.Amount)
	require.WithinDuration(t, entry.CreatedAt, output.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T){
	entry := createRandomEntry(t)
	updatedEntry := UpdateEntryParams {
		ID: entry.ID,
		Amount: util.RandomMoney(),
	}
	output, err := testQueries.UpdateEntry(GetContext(), updatedEntry)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, entry.ID, output.ID)
	require.Equal(t, entry.AccountID, output.AccountID)
	require.Equal(t, updatedEntry.Amount, output.Amount)
	require.WithinDuration(t, entry.CreatedAt, output.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T){
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(GetContext(), entry.ID)
	require.NoError(t, err)
	output, err := testQueries.GetEntry(GetContext(), entry.ID)
	require.Error(t, err)
	require.Empty(t, output)
	
}

func TestListEntry(t *testing.T){
	for i:=0; i< 10;i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(GetContext(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

}