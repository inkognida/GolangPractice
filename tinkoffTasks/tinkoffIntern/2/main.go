package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Val int
	Op  string
}

func main() {

	var r *bufio.Reader

	file, _ := os.Open("input.txt")
	r = bufio.NewReader(file)

	ns, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(ns))

	nums := make([]Pair, n)

	for i := 0; i < n; i++ {
		opers, _ := r.ReadString('\n')
		operations := strings.Split(strings.TrimSpace(opers), " ")
		tmp, _ := strconv.Atoi(operations[0])
		nums[i].Val = tmp
		nums[i].Op = operations[1]
	}

	x := 0
	for _, p := range nums {
		if p.Val >= 0 {
			x = x + p.Val
		}
	}
	fmt.Println(nums, x)
}
