package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Parsing a string to int using Atoi
    i, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Parsed integer:", i)
    }

    // Parsing a string to int with base 10 using ParseInt
    i64, err := strconv.ParseInt("456", 10, 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Parsed int64:", i64)
    }
}
