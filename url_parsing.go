package main

import (
    "fmt"
    "net/url"
)

func main() {
    // Example URL to parse
    rawURL := "https://example.com:8080/path?name=JohnDoe&age=25#section"

    // Parsing the URL
    u, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Error parsing URL:", err)
        return
    }

    // Accessing different parts of the URL
    fmt.Println("Scheme:", u.Scheme)
    fmt.Println("Host:", u.Host)
    fmt.Println("Path:", u.Path)
    fmt.Println("Query:", u.RawQuery)
    fmt.Println("Fragment:", u.Fragment)

    // Accessing URL query parameters
    queryParams := u.Query()
    fmt.Println("Name:", queryParams.Get("name"))
    fmt.Println("Age:", queryParams.Get("age"))
}
