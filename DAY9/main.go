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
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	fmt.Printf("Added %.2f to the wallet. New balance: %.2f\n", amount, w.Balance)
}

func (w *Wallet) SendMoney(amount float64, receiver *Wallet) error {
	if w.Balance < amount {
		return errors.New("insufficient balance")
	}
	w.Balance -= amount
	receiver.Balance += amount
	fmt.Printf("Sent %.2f. New balance: %.2f\n", amount, w.Balance)
	return nil
}

func (w *Wallet) CheckBalance() {
	fmt.Printf("Wallet balance: %.2f\n", w.Balance)
}

// User struct embedding Wallet
type User struct {
	Name        string
	Email       string
	PhoneNumber string
	Wallet      Wallet   // Embedded Wallet struct
}

func CreateUser(name, email, phoneNumber string, initialBalance float64) *User {
	return &User{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Wallet:      Wallet{Balance: initialBalance}, // Initialize embedded Wallet
	}
}

func showallusers(users []*User) {
	for i, user := range users {
		fmt.Printf("%d. %s\n", i+1, user.Name)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	users := []*User{}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Create a new user")
		fmt.Println("2. Check balance")
		fmt.Println("3. Add money")
		fmt.Println("4. Send money")
		fmt.Println("5. Show all users" )
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter name for the new user: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter email for the new user: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Enter phone number for the new user: ")
			phoneNumber, _ := reader.ReadString('\n')
			phoneNumber = strings.TrimSpace(phoneNumber)

			fmt.Print("Enter initial balance for the new user: ")
			initialBalanceStr, _ := reader.ReadString('\n')
			initialBalanceStr = strings.TrimSpace(initialBalanceStr)
			initialBalance, err := strconv.ParseFloat(initialBalanceStr, 64)
			if err != nil {
				fmt.Println("Invalid balance amount")
				continue
			}

			newUser := CreateUser(name, email, phoneNumber, initialBalance)
			users = append(users, newUser)
			fmt.Printf("User %s created with initial balance %.2f\n", name, initialBalance)

		case "2":
			fmt.Println("Choose user to check balance:")
			for i, user := range users {
				fmt.Printf("%d. %s\n", i+1, user.Name)
			}
			userChoiceStr, _ := reader.ReadString('\n')
			userChoiceStr = strings.TrimSpace(userChoiceStr)
			userChoice, err := strconv.Atoi(userChoiceStr)
			if err != nil || userChoice < 1 || userChoice > len(users) {
				fmt.Println("Invalid choice")
				continue
			}

			users[userChoice-1].Wallet.CheckBalance()

		case "3":
			fmt.Println("Choose user to add money:")
			for i, user := range users {
				fmt.Printf("%d. %s\n", i+1, user.Name)
			}
			userChoiceStr, _ := reader.ReadString('\n')
			userChoiceStr = strings.TrimSpace(userChoiceStr)
			userChoice, err := strconv.Atoi(userChoiceStr)
			if err != nil || userChoice < 1 || userChoice > len(users) {
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
			users[userChoice-1].Wallet.AddMoney(amount)

		case "4":
			fmt.Print("Enter amount to send: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}

			fmt.Println("Choose user to send from:")
			for i, user := range users {
				fmt.Printf("%d. %s\n", i+1, user.Name)
			}
			senderChoiceStr, _ := reader.ReadString('\n')
			senderChoiceStr = strings.TrimSpace(senderChoiceStr)
			senderChoice, err := strconv.Atoi(senderChoiceStr)
			if err != nil || senderChoice < 1 || senderChoice > len(users) {
				fmt.Println("Invalid choice")
				continue
			}

			fmt.Println("Choose user to send to:")
			for i, user := range users {
				fmt.Printf("%d. %s\n", i+1, user.Name)
			}
			receiverChoiceStr, _ := reader.ReadString('\n')
			receiverChoiceStr = strings.TrimSpace(receiverChoiceStr)
			receiverChoice, err := strconv.Atoi(receiverChoiceStr)
			if err != nil || receiverChoice < 1 || receiverChoice > len(users) {
				fmt.Println("Invalid choice")
				continue
			}

			err = users[senderChoice-1].Wallet.SendMoney(amount, &users[receiverChoice-1].Wallet)
			if err != nil {
				fmt.Println(err)
			}

		case "5":
			showallusers(users)

		case "6":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
