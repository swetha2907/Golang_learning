package main
import (
    "fmt"
    "sync"
)

func main() {
    var counter int // the counter we want to increment
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
		fmt.Println("counter",counter)
        go func() {
            defer wg.Done()
            // Atomically increment the counter
            // atomic.AddInt64(&counter, 1)
			counter++
        }()
    }

    wg.Wait() // wait for all goroutines to finish
    fmt.Println("Final Counter:", counter)
}
