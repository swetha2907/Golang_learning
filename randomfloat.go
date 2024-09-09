package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println("Random float:", rand.Float64())
    fmt.Println("Random float:", rand.Float64())
}
