package main

import "fmt"

type name struct {
	r_no int
}

func main(){
	fmt.Println("hai")
	fmt.Println(name{203})
	fmt.Println(name{20603})

	fmt.Println(name{2063})
	s:=name{r_no:909}
	fmt.Println(s.r_no)
	s:=name{9091}
	fmt.Println(s.r_no)


}
