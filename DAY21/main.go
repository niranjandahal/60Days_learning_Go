package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		ch <- num
		fmt.Printf("Producer %d produced %d\n", id, num)
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Consumer %d consumed %d\n", id, num)
		time.Sleep(time.Millisecond * 500)
	}
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Worker %d processed %d\n", id, num)
		time.Sleep(time.Millisecond * 500)
	}
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fanOutProducer(ch chan<- int) {
	for i := 0; i < 15; i++ {
		num := rand.Intn(100)
		ch <- num
		fmt.Printf("Produced %d\n", num)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func main() {
	fmt.Println("Producer-Consumer Pattern")
	prodCh := make(chan int)
	var prodWg sync.WaitGroup
	var consWg sync.WaitGroup

	for i := 0; i < 3; i++ {
		prodWg.Add(1)
		go producer(prodCh, &prodWg, i+1)
	}

	for i := 0; i < 2; i++ {
		consWg.Add(1)
		go consumer(prodCh, &consWg, i+1)
	}

	prodWg.Wait()
	close(prodCh)
	consWg.Wait()

	fmt.Println("\nFan-In Pattern")
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			num := rand.Intn(100)
			ch1 <- num
			time.Sleep(time.Millisecond * 500)
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			num := rand.Intn(100)
			ch2 <- num
			time.Sleep(time.Millisecond * 500)
		}
		close(ch2)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			num := rand.Intn(100)
			ch3 <- num
			time.Sleep(time.Millisecond * 500)
		}
		close(ch3)
	}()

	out := fanIn(ch1, ch2, ch3)

	for num := range out {
		fmt.Printf("Received %d\n", num)
	}

	fmt.Println("\nFan-Out Pattern")
	fanOutCh := make(chan int)
	var fanOutWg sync.WaitGroup

	for i := 0; i < 3; i++ {
		fanOutWg.Add(1)
		go worker(i+1, fanOutCh, &fanOutWg)
	}

	go fanOutProducer(fanOutCh)

	fanOutWg.Wait()
}