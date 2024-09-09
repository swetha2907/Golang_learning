package main

import (
    "fmt"
    "strconv"
)

func main() {
    f, err := strconv.ParseFloat("3.14159", 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Parsed float:", f)
    }
}
