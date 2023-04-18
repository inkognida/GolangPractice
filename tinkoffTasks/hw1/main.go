package main

import (
	"fmt"
)

// SIMPLE CONST (NOT TASK)
const (
	size   = 11
	symbol = "X"
	color  = "\033[34m"
)

func sandglass() {
	for i := 0; i < size; i++ {
		fmt.Print(color + symbol)
	}

	fmt.Println()
	for i := 1; i < size-1; i++ {
		for j := 1; j < size; j++ {
			if j-1 == i || j == size-i {
				fmt.Print(color, symbol)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < size; i++ {
		fmt.Print(color + symbol)
	}
}

//FUNCTION TYPE
type params map[string]int
type paramInit func(args params)

func initSize(size int) paramInit {
	return func(args params) {
		args["size"] = size
	}
}

func initColor(color int) paramInit {
	return func(args params) {
		args["color"] = color
	}
}

func initSymbol(symbol int) paramInit {
	return func(args params) {
		args["symbol"] = symbol
	}
}

func sandglassV2(args ...paramInit) {
	defaultParams := params{"size": 15, "color": 2, "symbol": 'X'}

	for _, arg := range args {
		arg(defaultParams)
	}

	colorv2 := fmt.Sprintf("\033[3%dm", defaultParams["color"])
	symbolv2 := fmt.Sprintf("%c", defaultParams["symbol"])
	sizev2 := defaultParams["size"]

	for i := 0; i < sizev2; i++ {
		fmt.Print(colorv2 + symbolv2)
	}

	fmt.Println()
	for i := 1; i < sizev2-1; i++ {
		for j := 1; j < sizev2; j++ {
			if j-1 == i || j == sizev2-i {
				fmt.Print(colorv2, symbolv2)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < sizev2; i++ {
		fmt.Print(colorv2 + symbolv2)
	}
}

func main() {
	sandglass()
	fmt.Println()
	sandglassV2(initSize(15), initColor(5), initSymbol('@'))
}
