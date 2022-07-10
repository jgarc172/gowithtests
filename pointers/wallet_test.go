package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	expected := Bitcoin(0)
	got := wallet.Balance()
	if got != expected {
		t.Errorf("got '%v', expected '%v'", got, expected)
	}

	wallet.Deposit(Bitcoin(10))
	expected = expected + Bitcoin(10)
	got = wallet.Balance()
	if got != expected {
		t.Errorf("got '%v', expected '%v'", got, expected)
	}

	wallet.Withdraw(Bitcoin(5))
	expected = expected - Bitcoin(5)
	got = wallet.Balance()
	if got != expected {
		t.Errorf("got '%v', expected '%v'", got, expected)
	}
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func ExampleWallet() {
	w := Wallet{}
	fmt.Println(w.Balance())
	w.Deposit(Bitcoin(10))
	fmt.Println(w.Balance())
	w.Withdraw(Bitcoin(5))
	fmt.Println(w.Balance())
	// Output:
	// 0
	// 10
	// 5
}
