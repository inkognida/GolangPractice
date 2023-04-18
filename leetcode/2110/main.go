package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	max := int64(0)
	c := int64(0)

	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			c += max + 1
			max++
		} else {
			max = 0
		}
	}

	return c + int64(len(prices))
}

func main() {
	fmt.Println(getDescentPeriods([]int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 4, 3, 10, 9, 8, 7}))
	//fmt.Println(getDescentPeriods([]int{3, 2, 1, 4}))
}
