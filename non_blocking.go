package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message:", msg)
    case signals <- true:
        fmt.Println("sent signal")
    default:
        fmt.Println("no activity")
    }
}
