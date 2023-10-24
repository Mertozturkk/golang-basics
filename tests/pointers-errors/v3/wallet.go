package v3

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // balance int yerine Bitcoin kullanildi. Go dili mevcut veri tipinden yeni bir veri tipi olusturmayi saglar.
}

/*
 *Wallet ile Wallet arasindaki fark:
 *Wallet cagirildigi yerdeki Wallet'in kopyasini degil, aynisini gonderir.
 Wallet ise cagirildigi yerdeki Wallet'in kopyasini gonderir.
*/

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Go da hatalari bir degere atayip kullanabiliyoruz

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
