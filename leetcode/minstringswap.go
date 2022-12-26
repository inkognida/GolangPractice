package main

import (
	"fmt"
	"math"
	"strings"
)

func minSwaps(s string) int {
	count := 0
	str := []rune(strings.ReplaceAll(s, "[]", ""))

	left := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 91 {
			left++

		} else if str[i] != 91 && left != len(str)/2 {
			left = -1
			break
		} else if str[i] != 91 && left == len(str)/2 {
			return count
		}
	}

	if left != -1 {
		for i := 0; i < len(str)/4; i++ {
			count++
		}

		return count
	}

	for {
		if strings.Contains(string(str), "[]") {
			str = []rune(strings.ReplaceAll(string(str), "[]", ""))
		} else {
			break
		}
	}

	for i := 0; i < len(str)/2; i++ {
		count++
	}

	if count%2 == 0 {
		return count / 2
	} else {
		return int(math.Ceil(float64(count) / 2))
	}
}

func main() {

	s := "[[[[]]]]"
	//s_ := strings.Repeat("[", 500000)
	//s__ := strings.Repeat("]", 500000)

	r := minSwaps(s)

	fmt.Println(r)
}

// len(str)/2

// [[[]]]
// [][]
// [[[[ [] ]]]]
// [[[ [] ]]] 1
// [ [] [] [] ]
// [ ] [ ] [] [ ] [ ] 2
// unb++
// less=0
//
// if swap and 2 pairs = succeed else go next
//
// [[]]
