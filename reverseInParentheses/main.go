package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "foo(bar(baz))blim"
	fmt.Println(reverseInParentheses(str))
}

func reverseInParentheses(inputString string) string {
	r, _ := regexp.Compile(`\(([^()]*)\)`)

	for {
		matchs := r.FindAllString(inputString, -1)
		for _, v := range matchs {
			rev := v
			rev = strings.Replace(rev, "(", "", -1)
			rev = strings.Replace(rev, ")", "", -1)
			rev = reverseString(rev)

			inputString = strings.Replace(inputString, v, rev, -1)
		}

		if !strings.Contains(inputString, "(") {
			break
		}
	}

	return inputString
}

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
