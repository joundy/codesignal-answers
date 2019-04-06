package main

import "fmt"

func main() {
	picture := []string{"abc", "ded"}

	fmt.Println(addBorder(picture))
}

func addBorder(picture []string) []string {
	var tbBorder string

	for range picture[0] {
		tbBorder += "*"
	}

	picture = append(picture, tbBorder, "")
	copy(picture[1:], picture[0:])
	picture[0] = tbBorder

	for i := range picture {
		picture[i] = "*" + picture[i] + "*"
	}

	return picture
}
