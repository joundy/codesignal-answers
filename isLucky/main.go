package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1230

	fmt.Println(isLucky(n))
}

func isLucky(n int) bool {

	nToStr := strconv.Itoa(n)

	var half1 int64
	var half2 int64

	if len(nToStr)%2 != 0 {
		return false
	}

	for i := 0; i < len(nToStr)/2; i++ {
		i1, _ := strconv.ParseInt(string(nToStr[i]), 0, 0)
		i2, _ := strconv.ParseInt(string(nToStr[i+len(nToStr)/2]), 0, 0)
		half1 += i1
		half2 += i2
	}

	return half1 == half2
}
