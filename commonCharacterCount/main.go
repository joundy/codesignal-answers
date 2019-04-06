package main

import "fmt"

func main() {
	s1 := "aabcc"
	s2 := "adcaa"

	fmt.Println(commonCharacterCount(s1, s2))
}

func commonCharacterCount(s1 string, s2 string) int {
	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	var count int

	for i := range s1Rune {
		for j := range s2Rune {
			if s1Rune[i] == s2Rune[j] {
				count++
				s2Rune = append(s2Rune[:j], s2Rune[j+1:]...)
				break
			}
		}

	}

	return count
}
