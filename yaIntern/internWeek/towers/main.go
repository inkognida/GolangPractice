//package main
//
//import (
//	"fmt"
//	"os"
//	"sort"
//)
//
//func ReadN(all []int, i, n int) {
//	if n == 0 {
//		return
//	}
//	if m, _ := Scan(&all[i]); m != 1 {
//		os.Exit(1)
//	}
//	ReadN(all, i+1, n-1)
//}
//
//func Scan(a *int) (int, error) {
//	return fmt.Scan(a)
//}
//
//func main() {
//	var n, m int
//	_, _ = fmt.Scan(&n, &m)
//
//	plsW := make([]int, n)
//	ReadN(plsW, 0, n)
//	for i := 0; i < len(plsW)/2; i++ {
//		j := len(plsW) - i - 1
//		plsW[i], plsW[j] = plsW[j], plsW[i]
//	}
//	platforms := make(map[int]int, 0)
//	max := 0
//	maxPlace := 0
//	for _, tmp := range plsW {
//		if _, ok := platforms[tmp]; !ok && tmp > max {
//			platforms[tmp-max]++
//			if tmp-max > maxPlace {
//				maxPlace = tmp-max
//			}
//			max = tmp
//		}
//	}
//	guards := make([]int, m)
//	ReadN(guards, 0, m)
//	sort.Ints(guards)
//
//	count := 0
//	min := 0
//	mx := 0
//	for _, v := range guards {
//		mx = maxPlace
//		if len(platforms) == 0 {
//			break
//		}
//		for k, _ := range platforms {
//			if v <= k && k <= mx {
//				min = k
//				mx = k
//			}
//		}
//		if _, ok := platforms[min]; ok && platforms[min] > 0 {
//			platforms[min]--
//			count++
//			if platforms[min] == 0 {
//				delete(platforms, min)
//			}
//		}
//	}
//	fmt.Println(count)
//}


// file input solution

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

	platforms := make(map[int]int, 0)
	max := 0
	maxPlace := 0
	for _, v := range plsW {
		tmp, _ := strconv.Atoi(v)
		if _, ok := platforms[tmp]; !ok && tmp > max {
			platforms[tmp-max]++
			if tmp-max > maxPlace {
				maxPlace = tmp-max
			}
			max = tmp
		}
	}
	gwl, _ := r.ReadString('\n')
	gw := SplitBySeparator(strings.TrimSpace(gwl), ' ')

	guards := make([]int, len(gw))
	for i, v := range gw {
		tmp, _ := strconv.Atoi(v)
		guards[i] = tmp
	}
	sort.Ints(guards)

	count := 0
	min := 0
	m := 0
	for _, v := range guards {
		m = maxPlace
		if len(platforms) == 0 || v > maxPlace {
			break
		}
		for k, _ := range platforms {
			if v <= k && k <= m {
				min = k
				m = k
			}
		}
		if platforms[min] > 0 {
			platforms[min]--
			count++
			if platforms[min] == 0 {
				delete(platforms, min)
			}
		}
	}
	fmt.Println(count)
}

