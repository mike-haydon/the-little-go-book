package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, error := strconv.Atoi(os.Args[1])
	if error != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}
}
