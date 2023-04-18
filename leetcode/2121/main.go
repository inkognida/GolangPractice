package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getSum(idxs []int, idx int) int {
	s := 0
	for i := 0; i < len(idxs); i++ {
		if idxs[i] != idx {
			s += abs(idx - idxs[i])
		}
	}
	return s
}

func getDistances(arr []int) []int64 {
	nums := make(map[int][]int, 0)
	for i, n := range arr {
		if _, ok := nums[n]; !ok {
			nums[n] = make([]int, 0)
			nums[n] = append(nums[n], i)
		} else {
			nums[n] = append(nums[n], i)
		}
	}
	fmt.Println(nums)

	arr2 := make([]int64, len(arr))
	for i, n := range arr {
		arr2[i] = int64(getSum(nums[n], i))
	}

	return arr2
}

func main() {
	fmt.Println(getDistances([]int{2, 1, 3, 1, 2, 3, 3}))
}
