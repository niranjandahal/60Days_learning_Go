package main

import (
	"testing"
)

func BenchmarkAddMoney(b *testing.B) {
	wallet := Wallet{}
	b.ResetTimer() // Reset timer before benchmarking

	for i := 0; i < b.N; i++ {
		wallet.AddMoney(10)
	}
}

func BenchmarkSendMoney(b *testing.B) {
	sender := Wallet{Balance: 100}
	receiver := Wallet{}
	b.ResetTimer() // Reset timer before benchmarking

	for i := 0; i < b.N; i++ {
		sender.SendMoney(50, &receiver)
	}
}
