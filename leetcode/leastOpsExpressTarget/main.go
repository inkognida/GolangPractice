package main

import (
	"fmt"
	"math"
)

// TODO release this algo

func leastOpsExpressTarget(x int, target int) int {

	solutions := make(map[int]int, 0)
	r := 0
	maxOperations := target*2 - 1 // 3-> 3/3 + 3/3 + 3/3
	for i := 0; i < maxOperations; i++ {
		solutions[i] = 1
	}

	/*
		if even -> do something
		if odd -> do something

	*/
	return r
}
func minOperators(x, target int) int {
	memo := make(map[[2]int]int) // memoization table

	var dp func(int, int) int
	dp = func(x, target int) int {
		if target == x {
			return 0 // base case
		}
		if ops, ok := memo[[2]int{x, target}]; ok {
			return ops // memoization
		}

		// try all possible combinations of operations
		minOps := math.MaxInt32
		for i := 1; i <= target/2; i++ {
			j := target - i
			for _, op := range []string{"+", "-", "*", "/"} {
				if op == "+" && i > j {
					continue // avoid duplicate expressions
				}
				if op == "-" && i == j {
					continue // avoid expressions with zero value
				}
				if op == "*" && i > j {
					continue // avoid duplicate expressions
				}
				if op == "/" && i == 0 {
					continue // avoid division by zero
				}
				if op == "/" && j%i != 0 {
					continue // avoid irrational results
				}
				if op == "/" && i == j {
					continue // avoid expressions with non-integer result
				}
				var ops int
				if op == "+" {
					ops = dp(x, i) + dp(x, j) + 1
				} else if op == "-" {
					ops = dp(x, i) + dp(x, j) + 1
				} else if op == "*" {
					ops = dp(x, i) + dp(x, j) + 1
				} else if op == "/" {
					ops = dp(x, i) + dp(x, j) + 1
				}
				minOps = min(minOps, ops)
			}
		}

		memo[[2]int{x, target}] = minOps // memoization
		return minOps
	}

	return dp(x, target)
}

func leastOpsExpressTarget1(x int, target int) int {
	memo := make(map[int]int)
	return dfs(x, target, memo)
}

func dfs(x, target int, memo map[int]int) int {
	if target == 1 {
		return ops(x, 1) // only need one operator to get 1
	}
	if val, ok := memo[target]; ok {
		return val // return cached result if available
	}
	pos := 1
	exp := x
	for exp < target { // find the closest power of x to target
		pos++
		exp *= x
	}
	if exp == target {
		memo[target] = pos - 1 // only need pos-1 multiplication operators
		return pos - 1
	}
	res := pos + dfs(x, target-exp/x, memo) // use pos multiplication operator
	if exp/x*(x-1) >= target {              // use pos-1 multiplication operator
		res = min(res, pos-1+dfs(x, target-exp/x*(x-1), memo))
	}
	memo[target] = res
	return res
}

func ops(x, target int) int { // calculate the number of operators needed to express target with x
	if target == 0 {
		return 0
	}
	if target < x {
		return min(target*2-1, (x-target)*2)
	}
	res := ops(x, target/x) + 1 // divide by x
	if target%x == 0 {
		return res
	}
	return min(res+target%x*2, res+(x-target%x)*2) // add or subtract x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(leastOpsExpressTarget1(5, 501))

}
