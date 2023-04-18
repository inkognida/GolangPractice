package main

import "fmt"

func countHousePlacements(n int) int {
	s0, s1 := 0, 1
	mod := int(1e9 + 7)

	for i := 0; i <= n; i++ {
		s0, s1 = s1, (s0+s1)%mod
	}

	return (s1 * s1) % mod
}

func main() {
	fmt.Println(countHousePlacements(2))
}
