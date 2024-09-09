package main
import "fmt"

type stack[T any] struct {
	items []T
}

func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) pop() T {
	i := s.items[len(s.items)-1]
	s.items =s.items[:len(s.items)-1]
	
	return i; 
}

func main() {
	a := stack[int]{}
	a.push(2)
	fmt.Println(a.pop())

}