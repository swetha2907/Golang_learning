package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "one"
    }()

    for {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received from ch1:", msg1)
            return
        case msg2 := <-ch2:
            fmt.Println("Received from ch2:", msg2)
            return
        default:
            fmt.Println("No messages received yet, doing other work...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
