package main

import "fmt"

func main() {

	inputArray := []string{"aba", "aa", "ad", "vcd"}

	fmt.Println(allLongestStrings(inputArray))

}

func allLongestStrings(inputArray []string) []string {
	var temp int

	var longestString []string

	for _, v := range inputArray {
		if temp < len(v) {
			temp = len(v)
		}
	}

	for _, v := range inputArray {
		if len(v) == temp {
			longestString = append(longestString, v)
		}
	}

	return longestString
}
