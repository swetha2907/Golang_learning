package main

import (
    "fmt"
    "os"
)

func main() {
    // Open or create a file with write permission
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close() // Ensure the file is closed after writing

    // Write a string to the file
    _, err = file.WriteString("Hello, Go! File writing example.\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("Data successfully written to the file.")
}
