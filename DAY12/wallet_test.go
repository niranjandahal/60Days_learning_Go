package main

import (
	"testing"
)

func TestAddMoney(t *testing.T){
	w := Wallet{}
	w.AddMoney(100)

	if w.Balance != 100 {
		t.Errorf("Expected balance to be 100, got %.2f", w.Balance)
	}
}

func TestSendMoney(t *testing.T){
	sender := Wallet{Balance : 1000}
	reciever := Wallet{Balance : 1}

	err := sender.SendMoney(500,&reciever)

	if err != nil {
		t.Errorf("UnExpected no error, got %v", err)
	}

	if sender.Balance != 500 {
		t.Errorf("Expected balance to be 500, got %.2f", sender.Balance)
	}

	if reciever.Balance != 501 {
		t.Errorf("Expected balance to be 501, got %.2f", reciever.Balance)
	}
}