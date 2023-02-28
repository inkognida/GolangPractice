package main //making package for standalone executable

import (
	"fmt" // importing a package
	"log"
	"runtime"
	"strings"
	"time"
)

// function block

func	changeFunc(value *int) {
	*value++
}

func	returnFunc (x, y int ) (x1, y1 int) {
	x1 = x * 10
	y1 = y * 10
	return
}

func	blankIdentifier() (a, b, c int) {
	return 1, 2, 3
}

func	variadicFunc(names ...string) {
	if len(names) == 0 {
		fmt.Println("No names!")
		return
	}
	for _, name := range names {
		fmt.Printf("%s ", name)
	}
	fmt.Println()
	return
}

func	calculationFunc(a, b int) (sum, product, diff int) {
	return a + b, a * b, a - b
}

func	sumOfSeries(values ...int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return
}

func	deferFuncMain() {
	fmt.Println("Main defer func")
	defer deferFuncHelp()
	fmt.Println("before DEFER done")

	for i := 0; i < 5; i++ { // possible leak
		defer fmt.Println("index: ", i)
	}
}

func	deferFuncHelp() {
	fmt.Println("Defer func help")
}

// tracing execution of functions, same with constructor/destructor in c++

func	track(constructor string) {
	fmt.Println("constructor", constructor)
}

func	untracked(destructor string) {
	fmt.Println("destructor", destructor)
}

func	childTrack() {
	track("child")
	defer untracked("child")
	fmt.Println("child is on")
}

func	parentTrack() {
	track("parent")
	defer untracked("parent")
	fmt.Println("parent is on, next step is to turn on the child")
	childTrack()
}

func	loggingDefer(name string) (value int, error_ error){
	defer func() {
		fmt.Printf("Name: %s, value: %d, error code: %v\n", name, value, error_)
	} ()
	return 10, nil
}

func	hierarchyFactorial(n uint64) (facSum uint64) {
	if n == 0 || n == 1 {
		return 1
	}
	facSum = 1
	for i := 1; uint64(i) <= n; i++ {
		facSum *= uint64(i)
	}
	return
}

func	asParameter(values ...int) {
	fmt.Println(values)
}

func	callParameter(name string, function func(...int)) {
	fmt.Printf("Name: %s ", name)
	function(1, 2, 3, 4)
}


// filter function

type filter func(int) bool // avoid func(int) bool inside sliceForming_func
type split	func([] int) ([] int, [] int)

func	isOdd(value int) bool {
	if value % 2 != 0 {
		return true
	}
	return false
}

func	isEven(value int) bool {
	if value % 2 == 0 {
		return true
	}
	return false
}

func	sliceForming(slice[] int, function filter)(result[] int) {
	for _, value := range slice {
		if function(value) {
			result = append(result, value)
		}
	}
	return
}

func	twoSliceForming(slice[] int, function filter) (odd, even[] int) {
	for _, value := range slice {
		if function(value) {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	return
}

func	assigningFunc() {
	printFunc := func(a, b int) int {return a + b}
	delta := 10

	for i := 1; i < 2; i++ {
		fmt.Println("Assigning function:", printFunc(i, delta))
	}
}

func	addDelta(a int) func(b int) int { // returning a function using closures
	return func(b int) int {
		return a + b
	}
}

func	adder() func(int) int {
	var value int // value is 0
	return func(delta int) int {
		value += delta
		return value
	}
}

func	fibonacciClosure() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a + b
		return b
	}
}

func	factoryFunction(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func	sliceFormingFactory(function filter) split {
	return func(slice[] int) (x, y[] int) {
		for _, value := range slice {
			if function(value) {
				x = append(x, value)
			} else {
				y = append(y, value)
			}
		}
		return
	}
}

// debugging block

func	debuggingRuntime() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s: %d", file, line)
	}
	fmt.Println("here")
	where()
	fmt.Println(10 + 10)
	where()
	fmt.Println("end of func")
	where()
}

func	debuggingLog() {
	log.SetFlags(log.Llongfile)
	var where = log.Print

	where()
	fmt.Println("Log is here")
	where()
}

func	timingFunc() {
	start := time.Now()
	for i := 0; i < 1000; i++ {}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("Time to execute:", diff)
}

func main(){ // entry point to my program
	funcValue := 10
	changeFunc(&funcValue)
	fmt.Println(funcValue)
	fmt.Println(returnFunc(10, 20))
	_, _, blankValue := blankIdentifier()
	fmt.Println(blankValue)
	variadicFunc("Mike", "John", "Kate")
	fmt.Println(calculationFunc(3, 4))
	fmt.Println(sumOfSeries(1, 2, 3, 4))

	deferFuncMain()
	parentTrack()
	loggingDefer("Tom")
	fmt.Println(hierarchyFactorial(10))
	callParameter("Nikki", asParameter)

	slice := [] int {1, 2, 3, 4, 5, 6, 7}
	even := sliceForming(slice, isEven)
	fmt.Println(even)
	odd := sliceForming(slice, isOdd)
	fmt.Println(odd)

	fmt.Println(twoSliceForming(slice, isEven))
	assigningFunc()

	addDelta_ := addDelta(10)
	fmt.Println("Return function example:", addDelta_(10))

	adder_ := adder() // function closure
	fmt.Println(adder_(1))
	fmt.Println(adder_(20))
	fmt.Println(adder_(300))

	fibonacciClosure_ := fibonacciClosure()
	for i := 2; i < 5; i++ {
		fmt.Println("Fibonacci closure:", i, "-", fibonacciClosure_())
	}

	addJpeg := factoryFunction(".jpeg") // add any suffix
	addJpeg_ := addJpeg("example")
	fmt.Println("Default factory function with adding a suffix:", addJpeg_)

	factorySlice_ := [] int {1, 2, 3, 4, 5}
	odd_, even_ := sliceFormingFactory(isEven)(factorySlice_)
	fmt.Println("sliceFormingFactory odd and even example:", odd_, even_)

	debuggingRuntime()
	debuggingLog()
	timingFunc()

}


