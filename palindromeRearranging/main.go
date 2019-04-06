package main

import (
	"fmt"
)

func main() {
	inputString := "abbcabb"
	fmt.Println(palindromeRearranging(inputString))
}

func palindromeRearranging(inputString string) bool {

	temp := make(map[string]int)
	var count int

	for _, v := range inputString {
		temp[string(v)]++
	}

	for _, v := range temp {
		if v%2 != 0 {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}
