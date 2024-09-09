package main

import "fmt"

// send function can only send data to the channel
func send(ch chan<- int, value int) {
    ch <- value
}

// receive function can only receive data from the channel
func receive(ch <-chan int) {
    value := <-ch
    fmt.Println("Received:", value)
}

func main() {
    ch := make(chan int)

    go send(ch, 42)
    receive(ch)
}
