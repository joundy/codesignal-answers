package main

import "fmt"

func main() {

	a := []int{50, 60, 60, 45, 70}

	fmt.Println(alternatingSums(a))
}

func alternatingSums(a []int) []int {

	var teamA []int
	var teamB []int

	for i, v := range a {
		if i%2 == 0 {
			teamA = append(teamA, v)
			continue
		}
		teamB = append(teamB, v)
	}

	fmt.Println(teamA)

	return []int{sumSliceInt(teamA), sumSliceInt(teamB)}
}

func sumSliceInt(v []int) int {
	var sum int

	for _, v := range v {
		sum += v
	}

	return sum
}
