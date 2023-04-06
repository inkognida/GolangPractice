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

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func main() {
	var r *bufio.Reader
	file, _ := os.Open("input.txt")
	r = bufio.NewReader(file)
	r.ReadString('\n')
	plsWline, _ := r.ReadString('\n')
	plsW := SplitBySeparator(strings.TrimSpace(plsWline), ' ')
	reverse(plsW)
	platforms := make(map[int]struct{}, 0)
	pls := make([]int, 0)
	max := 0
	for _, v := range plsW {
		tmp, _ := strconv.Atoi(v)
		if _, ok := platforms[tmp]; !ok && tmp > max {
			platforms[tmp] = struct{}{}
			pls = append(pls, tmp-max)
			max = tmp
		}
	}
	sort.Ints(pls)
	gwl, _ := r.ReadString('\n')
	gw := SplitBySeparator(strings.TrimSpace(gwl), ' ')

	l := len(pls)
	c := 0
	for _, v := range gw {
		tmp, _ := strconv.Atoi(v)
		if l > 0 {
			for j, _ := range pls {
				if tmp <= pls[j] {
					c++
					l--
					pls[j] = 0
					break
				}
			}
		} else {
			break
		}
	}
	fmt.Println(c)
}
