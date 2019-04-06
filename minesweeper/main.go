package main

import "fmt"

func main() {
	matrix := [][]bool{
		[]bool{true, false, false},
		[]bool{false, true, false},
		[]bool{false, false, false},
	}

	fmt.Println(minesweeper(matrix))
}

func minesweeper(matrix [][]bool) (result [][]int) {
	var neighbor [][]int
	neighbor = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for x := range matrix {
		var cur []int
		for y := range matrix[x] {
			var curSum int
			for _, v := range neighbor {
				res := getValue(matrix, x+v[0], y+v[1])
				curSum += res
			}
			cur = append(cur, curSum)
		}
		result = append(result, cur)
	}

	return
}

func getValue(matrix [][]bool, x, y int) int {
	if x < 0 || x > len(matrix)-1 {
		return 0
	}

	if y < 0 || y > len(matrix[x])-1 {
		return 0
	}

	if matrix[x][y] {
		return 1
	}

	return 0
}
