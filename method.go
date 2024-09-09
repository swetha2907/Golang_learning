package main
import "fmt"
type rect struct{
	height int
	width int 
}
func (r rect) peri() int {
    return 2*r.width + 2*r.height
}

func main() {
	r:=rect{height:20,width:10}

	fmt.Println(r.peri())
}
