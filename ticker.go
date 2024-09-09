package main
import (
    "fmt"
    "time"
)

func main() {
    // Create a new ticker that triggers every second
    ticker := time.NewTicker(1 * time.Second)
    done := make(chan bool)

    // Start a goroutine that will stop the ticker after 5 seconds
    go func() {
        time.Sleep(5 * time.Second)
        done <- true
    }()

    // Use the ticker to print the current time every second
    for {
        select {
        case t := <-ticker.C:
            fmt.Println("Tick at", t)
        case <-done:
            ticker.Stop()
            fmt.Println("Ticker stopped")
            return
        }
    }
}
