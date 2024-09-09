package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    fmt.Println("Default format:", now)

    // Custom formats using the reference time layout
    fmt.Println("Format 1:", now.Format("2006-01-02 15:04:05"))
    fmt.Println("Format 2:", now.Format("02-Jan-2006"))
    fmt.Println("Format 3:", now.Format("Monday, 02-Jan-06 03:04:05 PM"))
    fmt.Println("Format 4:", now.Format("2006-01-02 15:04:05 MST"))
}
