package main

import "fmt"

func main() {

	var matrix [][]int

	matrix = [][]int{
		[]int{0, 1, 1, 2},
		[]int{0, 5, 0, 0},
		[]int{2, 0, 3, 3},
	}

	fmt.Println(matrixElementsSum(matrix))
}

func matrixElementsSum(matrix [][]int) int {
	var count int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				count += matrix[i][j]
				continue
			}
			for z := 0; z < i; z++ {
				if matrix[z][j] == 0 {
					matrix[i][j] = 0
					continue
				}
			}
			count += matrix[i][j]
		}
	}
	return count
}
