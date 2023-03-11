package main

import (
	"fmt"
)

func maxWidthRamp(nums []int) int {
	var ii int
	r := 0

	j, nj := len(nums)-1, nums[len(nums)-1]
	i := 0
	for i < len(nums) {
		if nums[i] <= nj && i < j {
			ii = i
			if j-ii > r {
				r = j - ii
			} else {
				i = j
				continue
			}
		} else if i == j && j != 0 {
			j, nj = j-1, nums[j-1]
			i = 0
			continue
		}
		i++
	}

	return r
}

func main() {
	set := []int{6, 7, 8, 8, 6, 5, 5, 8, 2, 2}
	set2 := []int{2, 2, 1}

	fmt.Println(maxWidthRamp(set))
	fmt.Println(maxWidthRamp(set2))
}
