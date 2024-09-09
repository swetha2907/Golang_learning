package main

import (
    "fmt"
    "math/rand"
    // "time"
)

func main() {
    // Seed the random number generator with the current time
    // rand.Seed(time.Now().UnixNano())

    // Generate random integers between 0 and 99
    fmt.Println("Random int:", rand.Intn(100))
    fmt.Println("Random int:", rand.Intn(100))
    fmt.Println("Random int:", rand.Intn(100))
}
