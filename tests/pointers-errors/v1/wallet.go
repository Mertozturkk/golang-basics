package v1

import "fmt"

type Wallet struct {
	balance int // private acces modifier. Bu şekilde başlayan değişkenler sadece tanımlandıkları paket içerisinde kullanılabilirler.
}

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
