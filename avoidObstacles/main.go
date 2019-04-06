package main

import "fmt"

func main() {
	inputArray := []int{5, 3, 6, 7, 9}
	fmt.Println(avoidObstacles(inputArray))
}

func avoidObstacles(inputArray []int) (r int) {
	r = 1
Next:
	for {
		r++
		for _, v := range inputArray {
			if v%r == 0 {
				continue Next
			}
		}
		return r
	}
}
