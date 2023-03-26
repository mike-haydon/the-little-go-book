package main

import (
	"fmt"
	"strings"
)

func main() {
	// Arrays
	var scores [10]int
	scores[0] = 339

	scoresTwo := [4]int{9001, 9333, 212, 33}
	for index, value := range scoresTwo {
		fmt.Println("index", index, "value", value)
	}

	// Slices
	// Slice is not declared with a length
	//scores := []int{1, 4, 293, 4, 9}

	// Slice with length and capacity of 10
	//scores := make([]int, 10)

	// Slice with length of 0, but capacity of 10
	//scores := make([]int, 0, 10)

	// This will crash due to the slice not being expanded
	//scoreSlice := make([]int, 0, 10)
	//scoreSlice[7] = 9033
	//fmt.Println(scoreSlice)

	// Append to the end of a slice
	scoreSlice := make([]int, 0, 10)
	scoreSlice = append(scoreSlice, 5)
	fmt.Println(scoreSlice)

	// Re-slice slice
	scoreSliceTwo := make([]int, 0, 10)
	scoreSliceTwo = scoreSliceTwo[0:8]
	scoreSliceTwo[7] = 9033
	fmt.Println(scoreSliceTwo)

	// Add to slice with same size length and capacity
	scoreSliceThree := make([]int, 5)
	scoreSliceThree = append(scoreSliceThree, 9332)
	fmt.Println(scoreSliceThree)

	// Four common ways to initialize a slice
	//names := []string{"leto", "jessica", "paul"}
	//checks := make([]bool, 10)
	//var names []string
	//scores := make([]int, 0, 20)

	// Slices work on strings
	haystack := "the spice must flow"
	fmt.Println(strings.Index(haystack[5:], " "))

	// Go does not support negative values in slices
	// ex: to get all values except the last we do
	scoresLast := []int{1, 2, 3, 4, 5}
	scoresLast = scoresLast[:len(scoresLast)-1]
	fmt.Println(scoresLast)
}
