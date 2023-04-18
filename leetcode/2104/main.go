package main

import (
	"fmt"
	"math"
)

// All full version of algorithm
func All[T any](set []T) (subsets [][]T) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func minmaxdiff(n []int) int64 {
	max := math.MinInt
	min := math.MaxInt

	for _, v := range n {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	return int64(max - min)
}

func subArrayRanges(nums []int) int64 {
	r := int64(0)

	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			r += minmaxdiff(nums[i:j])
		}
	}

	return r
}

func main() {
	fmt.Println(subArrayRanges([]int{4, -2, -3, 4, 1}))
}
