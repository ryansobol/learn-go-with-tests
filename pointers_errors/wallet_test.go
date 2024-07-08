package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()

		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
	}

	t.Run("desposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}

		err := wallet.Withdraw(Bitcoin(10))

		assert.NoError(t, err)

		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		balanceStart := Bitcoin(20)

		wallet := Wallet{balanceStart}

		err := wallet.Withdraw(Bitcoin(21))

		assert.Equal(t, ErrInsufficientFunds, err)

		assertBalance(t, wallet, balanceStart)
	})
}
