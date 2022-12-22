package main

import (
	"fmt"
)

//func maxSatisfied(customers []int, grumpy []int, X int) int {
//	n := len(customers)
//	p1 := make([]int, n+1)
//	p2 := make([]int, n+1)
//	var g int
//	maxv := 0
//	for i := 0; i < n; i++ {
//		p1[i+1] = p1[i] + customers[i]
//		p2[i+1] = p2[i] + (1-grumpy[i])*customers[i]
//		if i+1 >= X {
//			g = (p1[i+1] - p2[i+1]) - (p1[i+1-X] - p2[i+1-X])
//			if maxv < g {
//				maxv = g
//			}
//		}
//	}
//	return p2[n] + maxv
//}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSatisfied_(customers []int, grumpy []int, minutes int) int {
	fmt.Println(customers, grumpy, minutes)

	len_ := len(customers)
	var direct_, secret_ int

	for i := 0; i < len_; i++ {
		if grumpy[i] == 0 {
			direct_ += customers[i]
			customers[i] = 0
		}
	}

	sum := 0
	for i, j := 0, 0; i < len_; i++ {
		sum += customers[i]
		if i - j == minutes {
			sum -= customers[j]
			j++
		}
		secret_ = maxInt(secret_, sum)
	}
	return secret_ + direct_
}

func main() {
	r := maxSatisfied_([]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1}, 3)

	fmt.Println(r)
}