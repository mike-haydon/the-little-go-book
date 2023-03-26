package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan int, 100)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: 1}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}
