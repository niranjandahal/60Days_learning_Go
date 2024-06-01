package main

import (
    "fmt"
)

func main() {
    fmt.Println("started learning golang")
    fmt.Print("Please enter your name: ")
    
    var name string
    fmt.Scanln(&name)

    if name != "" {
        fmt.Printf("Hello, %s! Nice to meet you.\n", name)
    } else {
        fmt.Println("Hello, World!")
    }
}
