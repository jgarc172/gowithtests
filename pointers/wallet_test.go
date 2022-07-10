package wallet

import (
	"fmt"
	"testing"
)

func ExampleWallet() {
	w := Wallet{}
	fmt.Println(w.Balance())
	w.Deposit(Bitcoin(10))
	fmt.Println(w.Balance())
	w.Withdraw(Bitcoin(5))
	fmt.Println(w.Balance())
	err := w.Withdraw(Bitcoin(10))
	fmt.Println(err)
	// Output:
	// 0
	// 10
	// 5
	// cannot withdraw, insufficient funds
}
func TestWallet(t *testing.T) {
	var wallet Wallet
	var expected Bitcoin
	var got Bitcoin
	var err error

	t.Run("new Wallet", func(t *testing.T) {
		wallet = Wallet{}
		expected = Bitcoin(0)
		got = wallet.Balance()
		if got != expected {
			t.Errorf("got '%v', expected '%v'", got, expected)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		wallet.Deposit(Bitcoin(10))
		expected = expected + Bitcoin(10)
		got = wallet.Balance()
		if got != expected {
			t.Errorf("got '%v', expected '%v'", got, expected)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		err = wallet.Withdraw(Bitcoin(5))
		if err != nil {
			t.Errorf("got '%v', expected '%v'", err, nil)
		}
		expected = expected - Bitcoin(5)
		got = wallet.Balance()
		if got != expected {
			t.Errorf("got '%v', expected '%v'", got, expected)
		}
	})

	t.Run("overdraft", func(t *testing.T) {
		err = wallet.Withdraw(Bitcoin(10))
		if err == nil {
			t.Errorf("got '%v', expected '%v'", err, "cannot withdraw, insufficient funds")
		}
	})
}
