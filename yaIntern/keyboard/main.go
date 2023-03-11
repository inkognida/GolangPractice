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

	file, err := os.Open("input.txt")
	if err == nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	sizeLine, _ := r.ReadString('\n')
	size, _ := strconv.Atoi(strings.TrimSpace(sizeLine))
	kb := make(map[string]string, size)

	keysLine, _ := r.ReadString('\n')
	keys := SplitBySeparator(strings.TrimSpace(keysLine), ' ')
	for i := 0; i < len(keys); i++ {
		kb[keys[i]] = ""
	}

	valuesLine, _ := r.ReadString('\n')
	values := SplitBySeparator(strings.TrimSpace(valuesLine), ' ')
	for i := 0; i < len(values); i++ {
		kb[keys[i]] = values[i]
	}

	_, _ = r.ReadString('\n')

	refLine, _ := r.ReadString('\n')
	ref := SplitBySeparator(strings.TrimSpace(refLine), ' ')
	res := 0
	for i := 1; i < len(ref); i++ {
		if kb[ref[i]] != kb[ref[i-1]] {
			res++
		}
	}
	fmt.Println(res)
}
