// package main

// import "fmt"

// func main() {
// 	type student struct {
// 		name string
// 		age int
// 	}
// 	func set (n string) *s{
// 		u :=s{name:n}
// 		u.age :=20
// 		return &u
// 	}
// 	fmt.Println(set("swetha"))
// }

package main

import "fmt"

type student struct {
	name string
	age  int
}

func set(n string) *student {
	u := student{name: n}
	u.age = 20
	return &u
}

func main() {
	fmt.Println(set("swetha"))

	d :=student{name:"abc", age:23}
	fmt.Println(d)

}
