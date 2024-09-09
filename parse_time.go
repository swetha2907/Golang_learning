package main

import (
    "fmt"
    "time"
)

func main() {
    layout := "2006-01-02 15:04:05"
    timeStr := "2024-08-20 10:34:45"

    parsedTime, err := time.Parse(layout, timeStr)
    if err != nil {
        fmt.Println("Error parsing time:", err)
    } else {
        fmt.Println("Parsed time:", parsedTime)
    }
}
