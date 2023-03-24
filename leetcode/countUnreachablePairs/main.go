package main

import (
	"fmt"
)

func dfs(node int, pairsMap map[int][]int, cnt *int64, visited map[int]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	*cnt++
	for _, pair := range pairsMap[node] {
		if visited[pair] == false {
			dfs(pair, pairsMap, cnt, visited)
		}
	}
}

func countPairs(n int, edges [][]int) int64 {
	pairsMap := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		pairsMap[i] = make([]int, 0)
	}

	for _, edge := range edges {
		pairsMap[edge[0]] = append(pairsMap[edge[0]], edge[1])
		pairsMap[edge[1]] = append(pairsMap[edge[1]], edge[0])
	}

	all := int64(n * (n - 1) / 2)
	visited := make(map[int]bool, n-1)
	for i := 0; i < n; i++ {
		if visited[i] == false {
			cnt := int64(0)
			dfs(i, pairsMap, &cnt, visited)
			all -= (cnt * (cnt - 1)) / 2
		}
	}

	return all
}

func main() {
	n := 50
	var edges [][]int
	edges = [][]int{
		{20, 0},
		{31, 1},
		{2, 15},
		{5, 31},
		{6, 44},
		{48, 7},
		{4, 8},
		{9, 3},
		{10, 44},
		{33, 11},
		{22, 12},
		{2, 13},
		{20, 14},
		{2, 16},
		{17, 25},
		{18, 22},
		{7, 19},
		{15, 21},
		{22, 23},
		{24, 11},
		{6, 26},
		{1, 27},
		{28, 2},
		{29, 6},
		{17, 30},
		{26, 32},
		{8, 34},
		{35, 5},
		{7, 37},
		{9, 38},
		{39, 36},
		{40, 20},
		{25, 41},
		{42, 41},
		{43, 14},
		{45, 33},
		{44, 46},
		{47, 41},
		{49, 2},
	}

	//fmt.Println(countPairs1(n, edges))
	fmt.Println(countPairs(n, edges))
	//fmt.Println("12" == "21")

	//for i := 0; i < n; i++ {
	//	for j := i + 1; j < n; j++ {
	//		fmt.Printf("(%d, %d) ", i, j)
	//	}
	//}
}
