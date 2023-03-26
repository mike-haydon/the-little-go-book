package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counterTwo = 0
	lock       sync.Mutex
)

func main() {
	for i := 0; i < 20; i++ {
		go incrTwo()
	}
	time.Sleep(time.Millisecond * 10)
}

func incrTwo() {
	lock.Lock()
	defer lock.Unlock()
	counterTwo++
	fmt.Println(counterTwo)
}
