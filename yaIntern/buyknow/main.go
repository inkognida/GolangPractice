package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SplitBySeparator(s string, sep rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
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
	for i, v := range p {
		pv, _ := strconv.Atoi(v)
		prices[i] = pv
	}

	md1 := 0
	md2 := 0

	st := make([][]int, 2)
	st[0] = make([]int, 4)
	st[1] = make([]int, 4)
	for i := 0; i < ds-3; i++ {
		for j := i + 1; j < ds-2; j++ {
			for k := j + 1; k < ds-1; k++ {
				for l := k + 1; l < ds; l++ {
					d1 := prices[j] - prices[i]
					d2 := prices[l] - prices[k]
					if d1 > md1 && i < j && j < k && k < l {
						md1 = d1
						st[0][0] = prices[i]
						st[0][1] = prices[j]

						st[1][0] = i + 1
						st[1][1] = j + 1
					}
					if d2 > md2 && i < j && j < k && k < l {
						md2 = d2
						st[0][2] = prices[k]
						st[0][3] = prices[l]

						st[1][2] = k + 1
						st[1][3] = l + 1
					}
				}
			}
		}
	}
	r0 := float64(1)
	b1, b1i := 0, 0
	s1, s1i := 0, 0
	if st[0][1]-st[0][0] > st[0][3]-st[0][2] {
		b1, b1i = st[0][0], st[1][0]
		s1, s1i = st[0][1], st[1][1]
	} else {
		b1, b1i = st[0][2], st[1][2]
		s1, s1i = st[0][3], st[1][3]
	}

	r1 := float64(1) / float64(b1) * float64(s1)
	r2 := r0 / float64(st[0][0]) * float64(st[0][1]) / float64(st[0][2]) * float64(st[0][3])

	if math.IsNaN(r1) {
		r1 = float64(0)
	} else if math.IsNaN(r2) {
		r2 = float64(0)
	}
	max := math.Max(math.Max(r0, r1), r2)
	//fmt.Println(r0, r1, r2, max, st)

	switch max {

	case r2:
		fmt.Println(2)
		fmt.Println(st[1][0], st[1][1])
		fmt.Println(st[1][2], st[1][3])
	case r1:
		fmt.Println(1)
		for i := 0; i < 3; i++ {
			if st[1][i] != 0 {
				fmt.Println(b1i, s1i)
			}
		}
	default:
		fmt.Println(0)
	}
}
