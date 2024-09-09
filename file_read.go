package main

import (
    "fmt"
    "io/ioutil"
    // "os"
)

func main() {
    // Reading an entire file into a string
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File content:\n", string(data))
}
