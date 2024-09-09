package main

import (
	"fmt"
)

type counterCmd struct {
	action string
	resp   chan int
}

func counter(state int, commands chan counterCmd) {
	for cmd := range commands {
		switch cmd.action {
		case "increment":
			state++
		case "decrement":
			state--
		case "get":
			cmd.resp <- state
		}
	}
}

func main() {
	commands := make(chan counterCmd)
	go counter(0, commands)

	// Increment the counter
	commands <- counterCmd{action: "increment"}

	// Decrement the counter
	commands <- counterCmd{action: "decrement"}

	// Get the current value of the counter
	response := make(chan int)
	commands <- counterCmd{action: "get", resp: response}
	fmt.Println("Counter value:", <-response)
}
