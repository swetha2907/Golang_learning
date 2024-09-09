// package main

// import (
//     "fmt"
// )

// func main() {
//     ch := make(chan int, 2)

//     ch <- 1
//     ch <- 2
// 	ch <- 3

//     fmt.Println(<-ch) 
//     fmt.Println(<-ch) 
//     fmt.Println(<-ch) 
// }

package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int, 2)

    go func() {
        ch <- 1
        fmt.Println("Sent 1")
        ch <- 2
        fmt.Println("Sent 2")
        ch <- 3
        fmt.Println("Sent 3")
    }()

    time.Sleep(1 * time.Second)
	fmt.Println("sleep")

    fmt.Println(<-ch) // Output: 1
    fmt.Println(<-ch) // Output: 2
    fmt.Println(<-ch) // Output: 3
}
