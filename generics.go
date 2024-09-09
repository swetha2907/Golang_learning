package main
import "fmt"

type generics interface {
	int | float64
}

func maximum[G generics](a G,b G) G {
	if(a>b) {
		return a;
	} else {
		return b;
	}
}
func printany [T any] (value T){ 
	println(value)  // any type
}
func main() {
	fmt.Println("aa")
	fmt.Println(maximum(2,3));
	printany("hi")
	
}

// package main

// import "fmt"

// // Define a generic type constraint
// type Number interface {
//     int | float64
// }

// // Generic function to find the maximum of two values
// func maximum[G Number](a, b G) G {
//     if a > b {
//         return a
//     }
//     return b
// }

// func main() {
//     fmt.Println("aa")
//     fmt.Println(maximum(2, 3))       // Output: 3
//     fmt.Println(maximum(3.5, 2.8))   // Output: 3.5
// }
