package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	Title     string
	Author    string
	Available bool
}

type Library struct {
	Books map[string]Book
}

// add book
func (l Library) AddBook(title, author string) Library {
	l.Books[title] = Book{title, author, true}
	fmt.Printf("Book %s added successfully\n", title)
	return l
}

//borrow book
func (l Library) BorrowBook(title string) Library {
	if l.Books[title].Available {
		l.Books[title] = Book{title,l.Books[title].Author,false}
		fmt.Printf("Book %s borrowed successfully\n", title)
	} else {
		fmt.Printf("Book %s is not available\n", title)
	}
	return l

}

//return book
func (l Library) ReturnBook(title string) Library {
	 if l.Books[title].Available {
		fmt.Printf("!!!Book %s is already available\n", title)
	 } else {
		l.Books[title] = Book{title,l.Books[title].Author,true}
		fmt.Printf("Book %s returned successfully\n", title)
	 }
	 return l
}

//display all books
func (l Library) Displayallbooks() {
	for key,value := range l.Books  {
		fmt.Printf("%-20s %-20s %-20t\n", key, value.Author, value.Available)
	}
}

func main(){

	library := Library{Books: make(map[string]Book)}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. Display All Books")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()
			library = library.AddBook(title, author)
		case "2":
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()
			library = library.BorrowBook(title)
		case "3":
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()
			library = library.ReturnBook(title)
		case "4":
			fmt.Println("\n----------------------------------------------------")
			fmt.Println("Title               Author              Available")
			library.Displayallbooks()
			fmt.Println("\n----------------------------------------------------")
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}

	}

}