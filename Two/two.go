package main

import "fmt"

type SaiyanTwo struct {
	Name  string
	Power int
}

func (s *SaiyanTwo) Super() {
	s.Power += 10000
}

func main() {
	goku := &SaiyanTwo{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power) // will print 19001
}
