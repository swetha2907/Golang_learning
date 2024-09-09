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
    p := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john@example.com",
    }

    // Marshal the Person struct to XML
    xmlData, err := xml.MarshalIndent(p, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(xmlData))
}
