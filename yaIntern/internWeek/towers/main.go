package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SplitBySeparator(s string, sep rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
}

func lds(arr []int) []int {
	n := len(arr)
	lds := make([]int, n)
	for i := range lds {
		lds[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] && lds[i] < lds[j]+1 {
				lds[i] = lds[j] + 1
			}
		}
	}
	max := 0
	end := 0
	for i := 0; i < n; i++ {
		if max < lds[i] {
			max = lds[i]
			end = i
		}
	}
	result := make([]int, max)
	result[max-1] = arr[end]
	for i := max - 2; i >= 0; i-- {
		for j := end - 1; j >= 0; j-- {
			if arr[j] > arr[end] && lds[j] == i+1 {
				result[i] = arr[j]
				end = j
				break
			}
		}
	}
	return result
}

func main() {
	var r *bufio.Reader

	if file, err := os.Open("input.txt"); err != nil {
		r = bufio.NewReader(os.Stdin)
	} else {
		r = bufio.NewReader(file)
	}

	dsLine, _ := r.ReadString('\n')
	d := SplitBySeparator(strings.TrimSpace(dsLine), ' ')
	n, _ := strconv.Atoi(d[0])
	m, _ := strconv.Atoi(d[1])

	plsWline, _ := r.ReadString('\n')
	plsW := SplitBySeparator(strings.TrimSpace(plsWline), ' ')
	pls := make([]int, n)
	for i, v := range plsW {
		tmp, _ := strconv.Atoi(v)
		pls[i] = tmp
	}

	gwl, _ := r.ReadString('\n')
	gw := SplitBySeparator(strings.TrimSpace(gwl), ' ')
	g := make([]int, m)
	for i, v := range gw {
		tmp, _ := strconv.Atoi(v)
		g[i] = tmp
	}

	platforms := lds(pls)
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j] // Reverse the order
	})

	for i := 0; i < len(platforms)-1; i++ {
		platforms[i] = platforms[i] - platforms[i+1]
	}
	platforms[len(platforms)-1] = platforms[len(platforms)-1] - 1

	c := 0
Main:
	for _, v := range g {
		for _, p := range platforms {
			if v <= p {
				c++
				continue Main
			}
		}
	}

	fmt.Println(c)
}
