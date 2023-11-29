package db

import (
	"testing"
	"time"

	"github.com/CarlosVinicius3258/go-simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer{
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg:= CreateTransferParams{
		FromAccountID: account1.ID,	
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
		}

	transfer, err := testQueries.CreateTransfer(GetContext(),arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
	
}

func TestGetTransfer(t *testing.T){
	transfer := createRandomTransfer(t)
	output, err := testQueries.GetTransfer(GetContext(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, transfer.ID, output.ID)
	require.Equal(t, transfer.FromAccountID, output.FromAccountID)
	require.Equal(t, transfer.ToAccountID, output.ToAccountID)
	require.Equal(t, transfer.Amount, output.Amount)
	require.WithinDuration(t, transfer.CreatedAt, output.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T){
	transfer := createRandomTransfer(t)
	updatedTransfer := UpdateTransferParams {
		ID: transfer.ID,
		Amount: util.RandomMoney(),
	}
	output, err := testQueries.UpdateTransfer(GetContext(), updatedTransfer)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, transfer.ID, output.ID)
	require.Equal(t, transfer.ToAccountID, output.ToAccountID)
	require.Equal(t, transfer.FromAccountID, output.FromAccountID)
	require.Equal(t, updatedTransfer.Amount, output.Amount)
	require.WithinDuration(t, transfer.CreatedAt, output.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T){
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(GetContext(), transfer.ID)
	require.NoError(t, err)
	output, err := testQueries.GetTransfer(GetContext(), transfer.ID)
	require.Error(t, err)
	require.Empty(t, output)
	
}

func TestListTransfer(t *testing.T){
	for i:=0; i< 10;i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(GetContext(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

}