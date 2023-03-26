package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
