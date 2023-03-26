package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)

	// returns 1
	total := len(lookup)
	fmt.Println(total)

	// has no return, can be called on a non-existing key
	delete(lookup, "goku")
	fmt.Println(len(lookup))

	// Set initial size of map
	//lookupTwo := make(map[string]int, 100)

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["kriilin"] = &Saiyan{Name: "Krillin"}
	fmt.Println(goku.Friends)

	lookupThree := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	for key, value := range lookupThree {
		fmt.Println(key, value)
	}
}

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}
