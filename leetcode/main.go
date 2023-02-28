package main

import (
	"fmt"
	"strings"
)

type lol int
type kek = lol // encapsulation methods from lol because of '='

func (arg *lol) Add(n int) {
	*arg += lol(n)
}

type Mod func(s string) string

func upper(s string) string {
	return strings.ToUpper(s)
}

func doubler(s string) string {
	return s + s
}

func applyer(s string, mods ...Mod) string {
	for _, mod := range mods {
		s = mod(s)
	}

	return s
}

func main() {
	a := make([]string, 100)
	//a = append(a, "hello")
	println(len(a), cap(a))

	b := make([]int, 4, 4)
	b = append(b, 123)
	b = append(b, 123)
	println(len(b), cap(b), "slice:", b)

	fmt.Println(b)

	c := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		c[i] = i
	}
	c1 := c[2:4]
	c1 = append(c1, 123)
	c1 = append(c1, 123)
	c = append(c, 222)

	fmt.Println(c, c1)

	var ex lol
	ex = 10
	ex.Add(5)

	var sas kek
	sas.Add(123)

	println(ex, sas)

	println(applyer("Go", upper, doubler))
}


package main

import "fmt"

type SandglassArgMap map[string]int
type Sandglass func(args SandglassArgMap)

func sandglass(args ...Sandglass) {
	defaultArgs := SandglassArgMap{"size": 15, "color": 0, "char": 'X'}
	for _, arg := range args {
		arg(defaultArgs)
	}
	size := defaultArgs["size"]
	color := defaultArgs["color"]
	char := defaultArgs["char"]
	fmt.Println()
	for i := 0; i < size; i++ {
		fmt.Printf("\033[%dm%c\033[0m", color, char)
	}
	fmt.Println()
	for j := 1; j < size-1; j++ {
		for i := 0; i < size; i++ {
			if i == j || i == size-j-1 {
				fmt.Printf("\u001B[%dm%c\u001B[0m", color, char)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < size; i++ {
		fmt.Printf("\033[%dm%c\033[0m", color, char)
	}
	fmt.Println()
}

func getSandglassSize(size int) Sandglass {
	return func(args SandglassArgMap) {
		args["size"] = size
	}
}

func getSandglassColor(color int) Sandglass {
	return func(args SandglassArgMap) {
		args["color"] = color
	}
}

func getSandglassChar(char int) Sandglass {
	return func(args SandglassArgMap) {
		args["char"] = char
	}
}

func main() {
	sandglass()
	sandglass(getSandglassChar('@'), getSandglassColor(33))
	sandglass(getSandglassSize(16))
}