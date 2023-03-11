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
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var r *bufio.Reader

	file, err := os.Open("input.txt")
	if err == nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	rp, _ := r.ReadString('\n')

	p := SplitBySeparator(strings.TrimSpace(rp), ' ')
	size, _ := strconv.Atoi(p[0])
	t, _ := strconv.Atoi(p[2])
	ideal, _ := strconv.Atoi(p[1])

	vLine, _ := r.ReadString('\n')
	vs := SplitBySeparator(strings.TrimSpace(vLine), ' ')

	vals := make([][]int, size)
	for i, v := range vs {

		n, _ := strconv.Atoi(v)
		vals[i] = make([]int, 2)
		vals[i][0] = n
		vals[i][1] = i + 1
	}

	cmp := func(i, j int) bool {
		return abs(vals[i][0]-ideal) < abs(vals[j][0]-ideal)
	}
	sort.Slice(vals, cmp)

	tmp := make([]int, 0)
	for _, v := range vals {
		if v[0] == ideal {
			tmp = append(tmp, v[1])
			continue
		}
		if t >= abs(v[0]-ideal) {
			tmp = append(tmp, v[1])
			t = t - abs(v[0]-ideal)
		} else {
			break
		}
	}
	fmt.Println(len(tmp))
	for i, element := range tmp {
		if i == len(tmp)-1 {
			fmt.Print(element, "\n")
			break
		}
		fmt.Print(element, " ")
	}
}
