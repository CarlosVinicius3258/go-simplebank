package db

import (
	"context"
	"testing"
	"time"

	"github.com/CarlosVinicius3258/simplebank/util"
	"github.com/stretchr/testify/require"
)
func getContext() context.Context{
	return context.Background()
}
func createRandomAccount(t *testing.T) Account{
	arg:= CreateAccountParams{
			Owner: util.RandomOwner(), // randomly generated?
			Balance: util.RandomMoney(),
			Currency: util.RandomCurrency(),
		}

	account, err := testQueries.CreateAccount(getContext(),arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
	
}

func TestGetAccount(t *testing.T){
	account := createRandomAccount(t)
	output, err := testQueries.GetAccount(getContext(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, account.ID, output.ID)
	require.Equal(t, account.Owner, output.Owner)
	require.Equal(t, account.Balance, output.Balance)
	require.Equal(t, account.Currency, output.Currency)
	require.WithinDuration(t, account.CreatedAt, output.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	account := createRandomAccount(t)
	updatedAccount := UpdateAccountParams {
		ID: account.ID,
		Balance: util.RandomMoney(),
	}
	output, err := testQueries.UpdateAccount(getContext(), updatedAccount)
	require.NoError(t, err)
	require.NotEmpty(t, output)
	require.Equal(t, account.ID, output.ID)
	require.Equal(t, account.Owner, output.Owner)
	require.Equal(t, updatedAccount.Balance, output.Balance)
	require.Equal(t, account.Currency, output.Currency)
	require.WithinDuration(t, account.CreatedAt, output.CreatedAt, time.Second)



}

func TestDeleteAccount(t *testing.T){
	account:= createRandomAccount(t)
	err := testQueries.DeleteAccount(getContext(), account.ID)
	require.NoError(t, err)
	output, err := testQueries.GetAccount(getContext(), account.ID)
	require.Error(t, err)
	require.Empty(t, output)
	
}

func TestListAccount(t *testing.T){
	for i:=0; i< 10;i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(getContext(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

}