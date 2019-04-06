package main

import "fmt"

func main() {
	image := [][]int{
		[]int{7, 4, 0, 1},
		[]int{5, 6, 2, 2},
		[]int{6, 10, 7, 8},
		[]int{1, 4, 2, 0},
	}

	fmt.Println(boxBlur(image))
}

func boxBlur(image [][]int) (result [][]int) {

	for i := 1; i < len(image)-1; i++ {
		var cur []int
		for j := 1; j < len(image[i])-1; j++ {
			// blur
			b := image[i-1][j-1] + image[i-1][j] + image[i-1][j+1] + image[i][j-1] + image[i][j] + image[i][j+1] + image[i+1][j-1] + image[i+1][j] + image[i+1][j+1]

			cur = append(cur, b/9)
		}
		result = append(result, cur)
	}

	return
}
