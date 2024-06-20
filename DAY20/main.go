package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
    fmt.Println("Task started")

    for {
        select {
        case <-ctx.Done(): 
            fmt.Println("Task cancelled")
            return
        default:
            fmt.Print(".")
            time.Sleep(100 * time.Millisecond) // Simulate work by sleeping for 100 milliseconds
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() 

    go longRunningTask(ctx)

    <-ctx.Done()

    fmt.Println("\nMain function: context cancelled, exiting")
}
