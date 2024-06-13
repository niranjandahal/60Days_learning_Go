package main

import (
    "errors"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockWalletService is a mock for an external service that interacts with Wallet
type MockWalletService struct {
    mock.Mock
}

func (m *MockWalletService) AddMoney(w *Wallet, amount float64) {
    m.Called(w, amount)
    w.Balance += amount
}

func (m *MockWalletService) SendMoney(w *Wallet, amount float64, receiver *Wallet) error {
    args := m.Called(w, amount, receiver)
    if w.Balance < amount {
        return errors.New("insufficient balance")
    }
    w.Balance -= amount
    receiver.Balance += amount
    return args.Error(0)
}

func TestAddMoney(t *testing.T) {
    wallet := &Wallet{}
    service := new(MockWalletService)

    service.On("AddMoney", wallet, float64(100)).Return(nil)

    service.AddMoney(wallet, 100)

    service.Mock.AssertExpectations(t)
    assert.Equal(t, 100.0, wallet.Balance, "Expected balance to be 100")
}

func TestSendMoney(t *testing.T) {
    sender := &Wallet{Balance: 100}
    receiver := &Wallet{}
    service := new(MockWalletService)

    service.On("SendMoney", sender, float64(50), receiver).Return(nil)
    service.On("SendMoney", sender, float64(100), receiver).Return(errors.New("insufficient balance"))

    err := service.SendMoney(sender, 50, receiver)
    assert.NoError(t, err)
    assert.Equal(t, 50.0, sender.Balance, "Expected sender balance to be 50")
    assert.Equal(t, 50.0, receiver.Balance, "Expected receiver balance to be 50")

    err = service.SendMoney(sender, 100, receiver)
    assert.Error(t, err)
    assert.Equal(t, "insufficient balance", err.Error(), "Expected insufficient balance error")
}
