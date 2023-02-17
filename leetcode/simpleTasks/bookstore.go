package main

import (
	"fmt"
)

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
		if i-j == minutes {
			sum -= customers[j]
			j++
		}
		secret_ = maxInt(secret_, sum)
	}
	return secret_ + direct_
}

func main() {
	r := maxSatisfied_([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)

	fmt.Println(r)
}
