package main

import "fmt"

func main() {
	student := struct {
        name   string
        reg_no int
    }{
        "swetha",
        21,
    }
    fmt.Println(student)
}