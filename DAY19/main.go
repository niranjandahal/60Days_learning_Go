package main
import (
	"fmt"
	"sync"
)
type Counter struct {
    value int
    mu    sync.Mutex // mutex to protect the counter
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
    defer wg.Done()
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
    fmt.Printf("Counter value: %d\n", c.value)
}

func main() {
    counter := &Counter{value: 0}
    wg := &sync.WaitGroup{}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go counter.Increment(wg)
    }

    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.value)
}