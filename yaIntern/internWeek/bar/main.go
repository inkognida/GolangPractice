package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitBySeparator(s string, sep rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
}

func main() {
	var r *bufio.Reader

	if file, err := os.Open("input.txt"); err != nil {
		r = bufio.NewReader(os.Stdin)
	} else {
		r = bufio.NewReader(file)
	}

	dsLine, _ := r.ReadString('\n')
	d := SplitBySeparator(strings.TrimSpace(dsLine), ' ')
	n, _ := strconv.Atoi(d[0])
	m, _ := strconv.Atoi(d[1])

	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, m+1)
		tmp, _ := r.ReadString('\n')
		for j := 0; j < len(tmp)-1; j++ {
			matrix[i][j] = string(tmp[j])
		}
	}

	kl, _ := r.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kl))

	ingred := make([][]string, k)
	last := len(matrix) - 2
	for i := 0; i < k; i++ {
		ingred[i] = make([]string, 3)
		tmpIng, _ := r.ReadString('\n')
		ingred[i] = SplitBySeparator(tmpIng[:len(tmpIng)-1], ' ')
		count, _ := strconv.Atoi(ingred[i][1])
		for g := 0; g < count; g++ {
			for j, v := range matrix[last] {
				if v == " " {
					matrix[last][j] = ingred[i][2]
				}
			}
			last--
		}
	}

	for _, v := range matrix {
		fmt.Println(strings.Join(v[:], ""))
	}
}

// success
