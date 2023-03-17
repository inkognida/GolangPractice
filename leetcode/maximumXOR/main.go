package main

import "fmt"

//func maximumXOR(nums []int) int {
//	r := 0
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 && i > 0 {
//
//		}
//	}
//
//	return r
//}

func maximumXOR(nums []int) int {
	res := 0
	for _, i := range nums {
		res = res | i
	}
	return res
}

func main() {
	fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
	fmt.Println(0 | 3 | 2)
}

// x XOR (x AND y)
