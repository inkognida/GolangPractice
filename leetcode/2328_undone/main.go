package main

import "fmt"

func countPaths(grid [][]int) int {
	mod := int(1e9 + 7)
	r := len(grid) * (len(grid[0]))

	m := make(map[int]struct{}, 0)
	for _, v := range grid {
		for _, v1 := range v {
			m[v1] = struct{}{}
		}
	}
	fmt.Println(m, mod)

	return r
}

func reverse(nums *[]int) {
	f := 0
	l := len(*nums) - 1

	for f < l {
		tmp := (*nums)[f]
		(*nums)[f] = (*nums)[l]
		(*nums)[l] = tmp
		f++
		l--
	}
}

func main() {
	var grid [][]int

	grid = [][]int{{1, 1}, {3, 4}}

	arr := []int{3, 4, 5, 1, 8}
	reverse(&arr)
	fmt.Println(arr)
	fmt.Println(countPaths(grid))
}
