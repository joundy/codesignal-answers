package main

import "fmt"

func main() {
	inputArray := []int{1, 1, 1}

	fmt.Println(arrayChange(inputArray))
}

func arrayChange(inputArray []int) (result int) {

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i]-inputArray[i-1] < 1 {
			cur := inputArray[i-1] + 1 - inputArray[i]
			result += cur
			inputArray[i] = inputArray[i] + cur
		}
	}

	return
}
