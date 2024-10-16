package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

var numch = make(chan int, 100)

func producer(numCh chan<- int) {
	for {
		numCh <- rand.Intn(10)
		time.Sleep(time.Second * 1)
	}
}

func consumer(numCh <-chan int) {
	i := 0
	for num := range numCh {
		log.Printf("consumer: %d [%d]", num, i)
		i++
	}
}

func main() {
	pool, err := ants.NewPool(
		10,
		ants.WithPreAlloc(true),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Release()
	wg := sync.WaitGroup{}
	wg.Add(2)
	pool.Submit(func() {
		defer wg.Done()
		producer(numch)
	})
	pool.Submit(func() {
		defer wg.Done()
		consumer(numch)
	})

	wg.Wait()
}
