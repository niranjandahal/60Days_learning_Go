package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wallet struct {
	Owner   string
	Balance float64
}

func CreateWallet(owner string, initialBalance float64) *Wallet {
	return &Wallet{Owner: owner, Balance: initialBalance}
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	fmt.Printf("%s added %.2f to their wallet. New balance: %.2f\n", w.Owner, amount, w.Balance)
}

func (w *Wallet) SendMoney(amount float64, receiver *Wallet) error {
	if w.Balance < amount {
		return errors.New("insufficient balance")
	}
	w.Balance -= amount
	receiver.Balance += amount
	fmt.Printf("%s sent %.2f to %s. New balance: %.2f\n", w.Owner, amount, receiver.Owner, w.Balance)
	return nil
}

func (w *Wallet) CheckBalance() {
	fmt.Printf("%s's wallet balance: %.2f\n", w.Owner, w.Balance)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	wallets := []*Wallet{}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Create a new wallet")
		fmt.Println("2. Check balance")
		fmt.Println("3. Add money")
		fmt.Println("4. Send money")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter owner name for the new wallet: ")
			owner, _ := reader.ReadString('\n')
			owner = strings.TrimSpace(owner)

			fmt.Print("Enter initial balance for the new wallet: ")
			initialBalanceStr, _ := reader.ReadString('\n')
			initialBalanceStr = strings.TrimSpace(initialBalanceStr)
			initialBalance, err := strconv.ParseFloat(initialBalanceStr, 64)
			if err != nil {
				fmt.Println("Invalid balance amount")
				continue
			}

			newWallet := CreateWallet(owner, initialBalance)
			wallets = append(wallets, newWallet)
			fmt.Printf("Wallet for %s created with initial balance %.2f\n", owner, initialBalance)

		case "2":
			fmt.Println("Choose wallet to check balance:")
			for i, wallet := range wallets {
				fmt.Printf("%d. %s\n", i+1, wallet.Owner)
			}
			walletChoiceStr, _ := reader.ReadString('\n')
			walletChoiceStr = strings.TrimSpace(walletChoiceStr)
			walletChoice, err := strconv.Atoi(walletChoiceStr)
			if err != nil || walletChoice < 1 || walletChoice > len(wallets) {
				fmt.Println("Invalid choice")
				continue
			}

			wallets[walletChoice-1].CheckBalance()

		case "3":
			fmt.Println("Choose wallet to add money:")
			for i, wallet := range wallets {
				fmt.Printf("%d. %s\n", i+1, wallet.Owner)
			}
			walletChoiceStr, _ := reader.ReadString('\n')
			walletChoiceStr = strings.TrimSpace(walletChoiceStr)
			walletChoice, err := strconv.Atoi(walletChoiceStr)
			if err != nil || walletChoice < 1 || walletChoice > len(wallets) {
				fmt.Println("Invalid choice")
				continue
			}

			fmt.Print("Enter amount to add: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}
			wallets[walletChoice-1].AddMoney(amount)

		case "4":
			fmt.Print("Enter amount to send: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}

			fmt.Println("Choose wallet to send from:")
			for i, wallet := range wallets {
				fmt.Printf("%d. %s\n", i+1, wallet.Owner)
			}
			senderChoiceStr, _ := reader.ReadString('\n')
			senderChoiceStr = strings.TrimSpace(senderChoiceStr)
			senderChoice, err := strconv.Atoi(senderChoiceStr)
			if err != nil || senderChoice < 1 || senderChoice > len(wallets) {
				fmt.Println("Invalid choice")
				continue
			}

			fmt.Println("Choose wallet to send to:")
			for i, wallet := range wallets {
				fmt.Printf("%d. %s\n", i+1, wallet.Owner)
			}
			receiverChoiceStr, _ := reader.ReadString('\n')
			receiverChoiceStr = strings.TrimSpace(receiverChoiceStr)
			receiverChoice, err := strconv.Atoi(receiverChoiceStr)
			if err != nil || receiverChoice < 1 || receiverChoice > len(wallets) {
				fmt.Println("Invalid choice")
				continue
			}

			err = wallets[senderChoice-1].SendMoney(amount, wallets[receiverChoice-1])
			if err != nil {
				fmt.Println(err)
			}

		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
