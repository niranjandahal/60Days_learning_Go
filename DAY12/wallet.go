package main

import (
	"errors"
)

type Wallet struct {
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
}

func (w *Wallet) SendMoney(amount float64, reciever *Wallet) error {

	if w.Balance < amount {
		return errors.New("insufficient balance")
	}
	w.Balance -= amount
	reciever.Balance += amount
	return nil
}



