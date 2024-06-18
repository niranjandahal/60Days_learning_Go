package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
	time.Sleep(time.Second * 1)
}

func main() {
	//SEQUENTIAL 
	start := time.Now()

	display("Process 1")
	display("Process 2")
	display("Process 3")

	seqTime := time.Since(start)

	//CONCURRENT
	start = time.Now()

	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	
	display("Process for main function")

	conTime := time.Since(start)


	fmt.Printf("Sequential Time taken: %s\n", seqTime)
	fmt.Printf("Concurrent Time taken: %s\n", conTime)
}
