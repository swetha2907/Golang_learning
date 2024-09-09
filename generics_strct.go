package main
import "fmt"

// type Str struct {
// 	a int|float64;
// }

// func add[T Str](items T) T {
// 	T.a :=itemsS
// 	return T;
// }
// func main() {
// 	s=Str();
// 	fmt.Println(s.add(12))
// }



type Number interface {
	int | float64
}
type struct1[T Number] struct {
	a T
}

func add[T Number](i struct1[T], g T) struct1[T] {
	i.a += g;
	return i;
}

func (i *struct1[T]) add1(g T){
	i.a += g
}
func main() {
	h := struct1[int]{a:20}
	fmt.Println(add(h,3).a)
	f := struct1[int]{a: 30}
	f.add1(8);
	fmt.Println(f.a);
}