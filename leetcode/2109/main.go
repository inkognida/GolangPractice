package main

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var str []string

	idx := 0
	for _, sp := range spaces {
		str = append(str, s[idx:sp])
		idx = sp
	}
	str = append(str, s[idx:len(s)])

	return strings.Join(str, " ")
}

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}
