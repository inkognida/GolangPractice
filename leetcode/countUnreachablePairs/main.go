package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findInt(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

//func countPairs(n int, edges [][]int) int64 {
//	pairsMap := make(map[int]map[int]struct{}, n*(n-1)/2)
//
//	sort.Slice(edges, func(i, j int) bool {
//		return edges[i][0] < edges[j][0]
//	})
//
//	for i := 0; i < n; i++ {
//		pairsMap[i] = make(map[int]struct{}, 0)
//	}
//
//	for i := 0; i < len(edges); i++ {
//		if _, ok := pairsMap[edges[i][0]]; ok {
//			pairsMap[edges[i][0]][edges[i][1]] = struct{}{}
//			pairsMap[edges[i][1]][edges[i][0]] = struct{}{}
//		}
//	}
//
//	var keys []int
//	for k := range pairsMap {
//		keys = append(keys, k)
//	}
//	sort.Ints(keys)
//
//	for _, key := range keys {
//		var keys2 []int
//		for k := range pairsMap[key] {
//			keys2 = append(keys2, k)
//		}
//		sort.Ints(keys2)
//		for _, key2 := range keys2 {
//			if _, ok := pairsMap[key2]; ok {
//				for k, _ := range pairsMap[key2] {
//					if _, ok := pairsMap[key][k]; !ok {
//						if k == key {
//							continue
//						}
//						pairsMap[key][k] = struct{}{}
//					}
//				}
//			}
//		}
//	}
//
//	pairsMapAll := make(map[string]struct{}, n*(n-1)/2)
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			pairsMapAll[strconv.Itoa(i)+strconv.Itoa(j)] = struct{}{}
//		}
//	}
//	fmt.Println(pairsMapAll)
//	fmt.Println(pairsMap)
//
//	for mk, v := range pairsMap {
//		for kw, _ := range v {
//			key := strconv.Itoa(mk) + strconv.Itoa(kw)
//			if _, ok := pairsMapAll[key]; ok {
//				delete(pairsMapAll, key)
//			}
//		}
//	}
//
//	return int64(len(pairsMapAll))
//}

func countPairs(n int, edges [][]int) int64 {
	//pairsMap := make(map[int]map[int]struct{}, n*(n-1)/2)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	for i := range edges {
		sort.Ints(edges[i])
	}

	//for i := 0; i < n; i++ {
	//	pairsMap[i] = make(map[int]struct{}, 0)
	//}

	pairsMap := make(map[int][]int, n*(n-1)/2)
	for i := 0; i < n; i++ {
		pairsMap[i] = make([]int, 0)
	}

	fmt.Println(edges)
	for i := 0; i < len(edges); i++ {
		pairsMap[edges[i][0]] = append(pairsMap[edges[i][0]], edges[i][1])
		for j := 0; j < len(edges); j++ {
			if edges[j][0] == edges[i][0] {
				pairsMap[edges[i][0]] = append(pairsMap[edges[i][0]], edges[j][1])
			}
		}
	}
	fmt.Println(pairsMap)

	//for i := 0; i < len(edges); i++ {
	//	if _, ok := pairsMap[edges[i][0]]; ok {
	//		pairsMap[edges[i][0]][edges[i][1]] = struct{}{}
	//		pairsMap[edges[i][1]][edges[i][0]] = struct{}{}
	//	}
	//}
	//
	//for mk, v := range pairsMap {
	//	for kw, _ := range v {
	//		if _, ok := pairsMap[kw]; ok {
	//			for k, _ := range pairsMap[kw] {
	//				if _, ok := v[k]; !ok {
	//					if k == mk {
	//						continue
	//					}
	//					v[k] = struct{}{}
	//				}
	//			}
	//		}
	//	}
	//}

	pairsMapAll := make(map[string]struct{}, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairsMapAll[strconv.Itoa(i)+strconv.Itoa(j)] = struct{}{}
		}
	}
	fmt.Println(len(pairsMap), len(pairsMapAll))
	for mk, v := range pairsMap {
		for kw, _ := range v {
			key := strconv.Itoa(mk) + strconv.Itoa(kw)
			if _, ok := pairsMapAll[key]; ok {
				delete(pairsMapAll, key)
			}
			if _, ok := pairsMapAll[reverseString(key)]; ok {
				delete(pairsMapAll, key)
			}
		}
	}

	return int64(len(pairsMapAll))
}

func reverseString(str string) string {
	// Convert the string to a slice of runes
	runes := []rune(str)

	// Reverse the slice using a loop
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice back to a string
	return string(runes)
}

func main() {
	n := 11
	var edges [][]int
	edges = [][]int{
		{5, 0},
		{1, 0},
		{10, 7},
		{9, 8},
		{7, 2},
		{1, 3},
		{0, 2},
		{8, 5},
		{4, 6},
		{4, 2},
	}
	fmt.Println(countPairs(n, edges))
	fmt.Println("12" == "21")
}
