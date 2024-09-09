package main
import (
    "os"
    "text/template"
)

func main() {
    tmpl, err := template.New("example").Parse("Hello, {{.Name}}!")
    if err != nil {
        panic(err)
    }

    data := struct {
        Name string
    }{
        Name: "John",
    }

    err = tmpl.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}
