package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
    Email string
}

func main() {
    p := Person{
        Name:  "John",
        Age:   30,
        Email: "john@example.com",
    }

    // Marshal the Person struct to JSON
    jsonData, err := json.Marshal(p)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
}
