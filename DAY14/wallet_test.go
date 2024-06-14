package main

import (
	"testing"
)

func TestAddMoney(t *testing.T) {
	w := Wallet{}
	w.AddMoney(100)

	if w.Balance != 100 {
		t.Errorf("Expected balance to be 100, got %.2f", w.Balance)
	}
}

func TestSendMoney(t *testing.T) {
	sender := Wallet{Balance: 1000}
	receiver := Wallet{}

	err := sender.SendMoney(500, &receiver)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sender.Balance != 500 {
		t.Errorf("Expected sender balance to be 500, got %.2f", sender.Balance)
	}

	if receiver.Balance != 500 {
		t.Errorf("Expected receiver balance to be 500, got %.2f", receiver.Balance)
	}

	err = sender.SendMoney(1000, &receiver)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
