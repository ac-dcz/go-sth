package main

import (
	"fmt"
	"math/rand"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {
	q := pq.NewWith(func(a, b any) int {
		ia := a.(int)
		ib := b.(int)
		if ia > ib {
			return 1
		} else if ia < ib {
			return -1
		} else {
			return 0
		}
	})

	for i := 0; i < 10; i++ {
		q.Enqueue(rand.Int() % 100)
	}
	fmt.Println(q.Values()...)
	fmt.Println("======================")
	fmt.Println(q.String())
	fmt.Println("======================")
	s, _ := q.MarshalJSON()
	fmt.Println(string(s))
	fmt.Println("======================")
	iter := q.Iterator()

	for iter.Next() {
		fmt.Println(iter.Value(), iter.Index())
	}
	fmt.Println("===========================")
	for i := 0; i < 10; i++ {
		val, _ := q.Dequeue()
		fmt.Println(val)
	}
}
