package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SplitBySeparator(s string, sep rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type countPrintFunc struct {
	calls int
}

func (f *countPrintFunc) Print(args ...interface{}) (n int, err error) {
	f.calls++
	return fmt.Print(args...)
}

func (f *countPrintFunc) Calls() int {
	return f.calls
}

func main() {
	var r *bufio.Reader

	file, err := os.Open("input.txt")
	if err == nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	rp, _ := r.ReadString('\n')
	p := SplitBySeparator(strings.TrimSpace(rp), ' ')
	t, _ := strconv.Atoi(p[2])
	ideal, _ := strconv.Atoi(p[1])

	intStr, _ := r.ReadString('\n')
	intStr = intStr[:len(intStr)-1]

	intArrStr := make([]string, 0)
	start := 0
	for i := 0; i < len(intStr); i++ {
		if intStr[i] == ' ' {
			intArrStr = append(intArrStr, intStr[start:i])
			start = i + 1
		}
	}
	intArrStr = append(intArrStr, intStr[start:])

	indexes := make(map[int]int, len(intArrStr))
	intArrConv := make([]int, len(intArrStr))
	for i, str := range intArrStr {
		num, _ := strconv.Atoi(str)
		intArrConv[i] = num
		indexes[num] = i + 1
	}

	cmp := func(i, j int) bool {
		return abs(intArrConv[i]-ideal) < abs(intArrConv[j]-ideal)
	}
	sort.Slice(intArrConv, cmp)

	var myPrint = countPrintFunc{}
	for _, v := range intArrConv {
		if v == ideal {
			myPrint.Print(indexes[v], " ")
			continue
		}

		for i := 0; i < t; i++ {
			if v+i == ideal || v-i == ideal {
				myPrint.Print(indexes[v], " ")
				t = t - i
				break
			}
		}
	}
	fmt.Println(myPrint.Calls())

}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"sort"
//"strconv"
//"strings"
//)
//
//func main() {
//	// считываем входные данные
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	params := strings.Split(scanner.Text(), " ")
//	n, _ := strconv.Atoi(params[0])
//	x, _ := strconv.Atoi(params[1])
//	t, _ := strconv.Atoi(params[2])
//
//	scanner.Scan()
//	sculpturesStr := strings.Split(scanner.Text(), " ")
//	sculptures := make([]int, n)
//	for i, s := range sculpturesStr {
//		sculptures[i], _ = strconv.Atoi(s)
//	}
//
//	// считаем время, необходимое для приведения каждой скульптуры к идеальному весу
//	times := make([]int, n)
//	for i, s := range sculptures {
//		times[i] = (x - s + 1) / 2
//	}
//
//	// сортируем скульптуры по возрастанию времени
//	sort.Ints(times)
//	sort.Ints(sculptures)
//
//	// выбираем скульптуры в порядке их возможности сделать идеальными
//	ans := 0
//	for i, t1 := range times {
//		if t1 <= t {
//			t -= t1
//			ans++
//		} else {
//			break
//		}
//	}
//
//	// выводим результат
//	fmt.Println(ans)
//	for i := 0; i < ans; i++ {
//		fmt.Print(i+1, " ")
//	}
//}
