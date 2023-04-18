package main

import (
	"fmt"
	"strings"
)

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	s = strings.ReplaceAll(s, "()", "$")

	return 0
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
