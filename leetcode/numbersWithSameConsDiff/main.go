package main

import (
	"fmt"
	"math"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// TODO analyze it more

func numsSameConsecDiff(n int, k int) []int {
	r := make([]int, 0)
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	var f func(n int, pow int)
	f = func(n int, pow int) {
		if pow == 0 {
			r = append(r, n)
			return
		}

		ld := n % 10 // last digit of number -> need to generate next

		if ld+k < 10 {
			f(n*10+ld+k, pow-1)
		}

		if k != 0 && ld-k >= 0 {
			f(n*10+ld-k, pow-1)
		}

	}

	for i := 1; i <= 9; i++ {
		f(i, n-1)
	}

	return r
}

func main() {
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(3, 7))

}
