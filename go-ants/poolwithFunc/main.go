package main

import (
	"fmt"
	"log"
	"sync"

	"math/rand"

	"github.com/panjf2000/ants/v2"
)

type Request struct {
	Name string
}

func main() {
	wg := sync.WaitGroup{}
	pool, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		if req, ok := i.(*Request); ok {
			fmt.Println(req.Name)
			wg.Done()
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Release()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		pool.Invoke(&Request{
			Name: fmt.Sprintf("Dcz %d", rand.Intn(10)),
		})
	}
	wg.Wait()
}
