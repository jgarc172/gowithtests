package wallet

import "testing"

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
