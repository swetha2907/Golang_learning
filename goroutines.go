package main

import (
    "fmt"
    "time"
)

// Function to print numbers with a delay
func printNumbers(name string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", name, i)
       // time.Sleep(100 * time.Millisecond) // Sleep to simulate work
    }
}

func main() {
    // Launch two goroutines
    go printNumbers("Goroutine 1")
    go printNumbers("Goroutine 2")

    // Sleep to allow the goroutines to complete their execution
    time.Sleep(1 * time.Second)
    fmt.Println("Main function done")
}
