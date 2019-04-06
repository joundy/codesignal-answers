package main

import "fmt"

func main() {
	a := []int{1, 3, 2}
	b := []int{1, 2, 3}

	fmt.Println(areSimilar(a, b))
}

func areSimilar(a []int, b []int) bool {

	var count []int

	for i := range a {
		if a[i] != b[i] {
			count = append(count, i)
		}
	}

	if len(count) == 0 || len(count) == 2 && a[count[0]] == b[count[1]] && b[count[0]] == a[count[1]] {
		return true
	}

	return false
}
