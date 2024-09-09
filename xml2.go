package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
    Email   string   `xml:"email"`
}

func main() {
    xmlData := `<person><name>John Doe</name><age>30</age><email>john@example.com</email></person>`

    var p Person
    err := xml.Unmarshal([]byte(xmlData), &p)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", p)
}
