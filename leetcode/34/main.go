package main

import (
	"fmt"
)

func binarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	idx := binarySearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	var f, l int
	for i := idx; i >= 0 && nums[i] == target; i-- {
		f = i
	}
	for i := idx; i < len(nums) && nums[i] == target; i++ {
		l = i
	}
	return []int{f, l}
}

func main() {
	fmt.Println(searchRange([]int{3, 3, 3, 3, 3, 3, 3}, 3))
}
