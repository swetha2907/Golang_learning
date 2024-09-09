package main

import (
    "os"
    "text/template"
)

func main() {
    tmpl, err := template.New("example").Parse(`Items: {{range .Items}}{{.}} {{end}}`)
    if err != nil {
        panic(err)
    }

    data := struct {
        Items []string
    }{
        Items: []string{"apple", "banana", "cherry"},
    }

    err = tmpl.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}
