package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 150, 190, 170, -1, -1, 160, 180}

	fmt.Println(sortByHeight(a))
}

func sortByHeight(a []int) []int {

	var rawPeople []int

	// first remove the tree
	for _, v := range a {
		if v != -1 {
			rawPeople = append(rawPeople, v)
		}
	}

	sort.Ints(rawPeople)

	for i, v := range a {
		if v == -1 {
			rawPeople = append(rawPeople, 0)
			copy(rawPeople[i+1:], rawPeople[i:])
			rawPeople[i] = -1
		}
	}

	return rawPeople
}
