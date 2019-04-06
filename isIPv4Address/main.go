package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputString := "172.1s6.254.1"

	fmt.Println(isIPv4Address(inputString))
}

func isIPv4Address(inputString string) bool {

	groups := strings.Split(inputString, ".")
	if len(groups) != 4 {
		return false
	}

	for _, v := range groups {
		toInt, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		if !(toInt >= 0 && toInt <= 255) {
			return false
		}

	}
	return true
}
