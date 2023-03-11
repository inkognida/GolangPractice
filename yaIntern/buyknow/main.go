package main

import (
	"bufio"
	"fmt"
	"math"
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

func maxDiff(a []int) int {
	max := 0
	for i := 0; i < len(a)-1; i++ {
		diff := abs(a[i+1] - a[i])
		if diff > max {
			max = diff
		}
	}
	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func compareSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
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
	ds, _ := strconv.Atoi(d[0])

	prLine, _ := r.ReadString('\n')
	p := SplitBySeparator(strings.TrimSpace(prLine), ' ')

	prices := make([]int, ds)

	j := 0
	in := make(map[int][][]int, len(p)-3)

	for i, v := range p {
		pv, _ := strconv.Atoi(v)
		prices[i] = pv

		if j < len(p)-3 {
			in[j+1] = make([][]int, 2)
			in[j+1][0] = make([]int, 4)
			in[j+1][0] = []int{i, i + 1, i + 2, i + 3}
			j++
		}

	}

	act := make([][]int, len(p)-3)

	fmt.Println(in)

	for i := range act {
		act[i] = make([]int, 4)
		copy(act[i], prices[i:i+4])

		in[i+1][1] = make([]int, 4)
		in[i+1][1] = act[i]
	}

	sort.Slice(act, func(i, j int) bool {
		diff1 := maxDiff(act[i])
		diff2 := maxDiff(act[j])
		return diff1 > diff2
	})

	active := act[0]

	s0 := float64(1)
	s11 := float64(1) / float64(active[0]) * float64(active[1])
	s12 := float64(1) / float64(active[2]) * float64(active[3])
	s2 := s11 / float64(active[2]) * float64(active[3])

	max := math.Max(math.Max(s0, s11), math.Max(s12, s2))

	switch max {
	case s0:
		fmt.Println(0)
	case s11:
		fmt.Println(1)
		for _, mv := range in {
			if compareSlices(mv[1], active) {
				fmt.Println(mv[0])
			}
		}

	case s12:
		fmt.Println(1)
		for _, mv := range in {
			if compareSlices(mv[1], active) {
				fmt.Println(mv[0])
			}
		}
	case s2:
		fmt.Println(2)
		for _, mv := range in {
			if compareSlices(mv[1], active) {
				fmt.Println(mv[0], mv[1])
			}
		}
	}

}
