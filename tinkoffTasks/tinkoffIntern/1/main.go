package main

import (
	"fmt"
)

// Reading
//var r *bufio.Reader
//
//file, _ := os.Open("input.txt")
//r = bufio.NewReader(file)
//
//ns, _ := r.ReadString('\n')
//n, _ := strconv.Atoi(strings.TrimSpace(ns))
//
//numss, _ := r.ReadString('\n')
//numss1 := strings.Split(strings.TrimSpace(numss), " ")
//nums := make([]int, n)
//for i, v := range numss1 {
//	tmp, _ := strconv.Atoi(v)
//	nums[i] = tmp
//}
//fmt.Println(nums)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&nums[i])
	}

	for {
		op := 0
		for i := 0; i <= len(nums)-3; i++ {
			if nums[i+2] > nums[i] {
				nums[i], nums[i+2] = nums[i+2], nums[i]
				op = 1
			}
		}
		if op != 1 {
			break
		}
	}

	op2 := 0
	for {
		tmp := op2
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				op2++
			}
		}
		if op2 == tmp {
			break
		}
	}

	fmt.Println(op2)
}
