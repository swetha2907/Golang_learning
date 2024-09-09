package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string
    Age   int
    Email string
}

func main() {
    jsonData := `{"Name":"John","Age":30,"Email":"john@example.com"}`

    var p Person

    // Unmarshal the JSON data into a Person struct
    err := json.Unmarshal([]byte(jsonData), &p)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", p)
}
