package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a channel to send jobs to
    jobs := make(chan int, 5)
    // Ticker that ticks every 200 milliseconds
    ticker := time.NewTicker(10000 * time.Millisecond)
    defer ticker.Stop()

    // Start a goroutine to send jobs to the channel
    go func() {
        for i := 1; i <= 10; i++ {
            jobs <- i
            fmt.Println("Sent job", i)
        }
        close(jobs)
    }()

    // Process jobs at a limited rate
    for job := range jobs {
        <-ticker.C // Wait for the ticker to tick
        fmt.Println("Processing job", job)
    }

    fmt.Println("All jobs processed")
}
