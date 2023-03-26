package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WorkerTwo struct {
	id int
}

func (w WorkerTwo) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &WorkerTwo{id: 1}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
			// optional code here
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
