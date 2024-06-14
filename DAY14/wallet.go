package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
}

func (w *Wallet) SendMoney(amount float64, receiver *Wallet) error {
	if w.Balance < amount {
		return errors.New("insufficient balance")
	}
	w.Balance -= amount
	receiver.Balance += amount
	return nil
}

func main() {
	wallet := Wallet{}
	wallet.AddMoney(100)
	fmt.Printf("Wallet balance: %.2f\n", wallet.Balance)
}
