package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

func main() {
	st := arraystack.New()
	for i := 0; i < 5; i++ {
		st.Push(i)
	}
	for !st.Empty() {
		val, _ := st.Pop()
		fmt.Println(val.(int), st.Size())
	}
}
