package main
import (
    "fmt"
)

func main() {
    fmt.Println("Start")
    
    defer fmt.Println("This is deferred")
    
    fmt.Println("End")
}
