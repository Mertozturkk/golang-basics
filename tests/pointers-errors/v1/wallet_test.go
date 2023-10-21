package v1

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/*
Burada wallet.balance değişkeninin adresini yazdırıyoruz. Bu adresin test fonksiyonunda yazdırılan adres ile aynı olup olmadığını kontrol ediyoruz.
Gorulecegi uzere bu adresler farkli. Deposit methodunu kullanirken wallet.balance değişkeninin kopyasını kullanıyoruz. Bu yüzden test fonksiyonunda
wallet.balance değişkeninin adresi ile Deposit methodunda yazdırılan adres farklı oluyor.
*/
