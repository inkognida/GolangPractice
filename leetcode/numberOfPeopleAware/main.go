package main

import "fmt"

//func peopleAwareOfSecret(n int, delay int, forget int) int {
//	r := 0
//
//	know := make(map[int]bool, 0)
//	know[0] = true
//	for i := n; i > 1; i-- {
//
//	}
//
//	return r
//}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	// shares holds the day in which some person will start revealing the secret
	var shares [][2]int
	// forgets holds the day in which some person will start forgetting
	// the secret.
	var forgets [][2]int
	shares = append(shares, [2]int{1 + delay, 1})
	forgets = append(forgets, [2]int{1 + forget, 1})
	// total is the total number of people who knows the secret
	total := 1
	// curr is the number of people who can currently reveal the secret to someone
	curr := 0
	mod := int(1e9 + 7)
	for i := 1; i <= n; i++ {
		for len(forgets) > 0 && forgets[0][0] == i {
			curr = curr - forgets[0][1]
			total = total - forgets[0][1]
			forgets = forgets[1:]
		}
		for len(shares) > 0 && shares[0][0] == i {
			curr = curr + shares[0][1]
			shares = shares[1:]
		}
		total = (total + curr) % mod
		if curr > 0 {
			shares = append(shares, [2]int{i + delay, curr % mod})
			forgets = append(forgets, [2]int{i + forget, curr % mod})
		}
	}
	return total
}

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
}
