package main
import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"banana", "apple", "cherry", "kiwi", "blueberry"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	fmt.Println("Strings sorted by length:", strs)
}
