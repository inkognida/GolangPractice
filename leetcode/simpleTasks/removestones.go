package main

import (
	"container/heap"
	"fmt"
)

// TIMEOUT SOLUTION
/*
func maxElementIndex(slice []int) int {
	indexMax := 0
	max := 0

	for i, v := range slice {
		if v > max {
			max = v
			indexMax = i
		}
	}

	return indexMax
}

func minStoneSum(piles []int, k int) int {
	sum := 0

	if len(piles) == 1 {
		for ; k > 0; k-- {
			piles[0] = int(math.Ceil(float64(piles[0]) / float64(2)))
		}
		return piles[0]
	}

	for ; k > 0; k-- {
		index := maxElementIndex(piles)
		piles[index] = int(math.Ceil(float64(piles[index]) / float64(2)))
	}

	for _, v := range piles {
		sum += v
	}

	return sum
}

func main() {
	result := minStoneSum([]int{4122, 9928, 3477, 9942}, 6)

	fmt.Println(result)
}
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	h := IntHeap(piles)
	heap.Init(&h)
	sum := 0

	for ; k > 0; k-- {
		h[0] -= h[0] / 2
		heap.Fix(&h, 0)
	}

	for _, v := range piles {
		sum += v
	}

	return sum

}

func main() {
	piles := []int{4, 5, 9}
	k := 2

	r := minStoneSum(piles, k)
	fmt.Println("Answer", r)

}
