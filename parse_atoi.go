package main
import (
    "fmt"
    "strconv"
)
func main() {
    e, err := strconv.Atoi("103")
	_, err1 := strconv.Atoi("abc")
    if err1 != nil {
        fmt.Println("Error:", err1)
    }
	if err != nil {
        fmt.Println("Error:", err)
    }
	fmt.Println(e)
}
