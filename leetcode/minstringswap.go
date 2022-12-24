package main

import (
	"fmt"
	"math"
	"strings"
)

//Time Limit Exceeded
//func minSwaps(s string) int {
//	count := 0
//	str := []rune(strings.ReplaceAll(s, "[]", ""))
//
//	for {
//		if strings.Contains(string(str), "[]") {
//			str = []rune(strings.ReplaceAll(string(str), "[]", ""))
//		} else {
//			break
//		}
//	}
//
//	for i := 0; i < len(str)/2; i++ {
//		count++
//	}
//
//	if count%2 == 0 {
//		return count / 2
//	} else {
//		return int(math.Ceil(float64(count) / 2))
//	}
//
//}

func minSwaps(s string) int {
	count := 0
	str := []rune(strings.ReplaceAll(s, "[]", ""))

	//for {
	//	if strings.Contains(string(str), "[]") {
	//		str = []rune(strings.ReplaceAll(string(str), "[]", ""))
	//	} else {
	//		break
	//	}
	//}

	var flag bool
	flag = false

	for {
		if strings.Contains(string(str), "[]") {
			flag = true
			for i := 0; i < len(str)-1; i++ {
				if str[i] == 91 && str[i+1] == 93 {
					str = append(str[0:i-1], str[0:i+1]...)
				}
			}
		} else {
			break
		}
	}

	fmt.Println(string(str))

	if flag == false {
		for i := 0; i < len(str)/2; i++ {
			count++
		}
	} else {
		for i := 0; i < len(str); i++ {
			count++
		}
	}

	if count%2 == 0 {
		return count / 2
	} else {
		return int(math.Ceil(float64(count) / 2))
	}
}

func main() {
	s := "]]][[["
	r := minSwaps(s)

	fmt.Println(r)
}

// ]]][[[
// ][][
//
// ]][[
// [][]
// []][][
// for if v == [
// less++
// else
// less--
// if les <0
// unb++
// less=0
//
// if swap and 2 pairs = succeed else go next
//
