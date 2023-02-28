package main

import (
	"bytes"
	"fmt"
)

// arrays block

func	sumOfArray(array *[5]int) (sum int) {
	for i := range array {
		sum += array[i]
	}
	return
}

func	defaultArray() {
	var arr_[5] int
	for i := range arr_ {
		arr_[i] = i * 2
		fmt.Println("arrays values and indexes:", arr_[i], i)
	}

	var arr__ = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr__)

	var indexArr = [5]string{3: "Bob", 4: "Ron"}
	fmt.Println("Array with current indexes:", indexArr)
	fmt.Println("Sum of array using pointer to pass an array:", sumOfArray(&arr_)) // doesn't make copy of array inside function
}

var fib[10] int64
func	iterativeFibonacci() [10]int64{
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < 10; i++ {
		fib[i] = fib[i - 1] + fib[i - 2]
	}
	return fib
}

// slices block

func	defaultSlice() {
	var	arr [5]int
	slice := arr[1:4]

	for i := range arr {
		arr[i] = i
	}
	fmt.Println("Len of array and slice", len(arr), "-", len(slice),
		"\nCap of array and slice", cap(arr), "-", cap(slice))
	slice = slice[0:4]
	fmt.Println("slice:", slice, "array:", arr)
}

func	sliceFunc(slice []int) {
	for i := range slice {
		slice[i] += i
	}
}

func	makeSlice() {
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = i * 100
	}
	fmt.Println(slice)
}

// bytes and slices

func	bufferBytes() {
	var	buff bytes.Buffer

	buff.WriteString("Hello, world!")
	fmt.Println("Buffer message:", buff.String())
}

func	fibSlice(term int) (sliceFib_ []int) {
	sliceFib_ = make([]int, term)
	sliceFib_[0] = 0
	sliceFib_[1] = 1
	for i := 2; i < term; i++ {
		sliceFib_[i] = sliceFib_[i - 1] + sliceFib_[i - 2]
	}
	return
}

func	reSlicing() {
	slice := make([]int, 0, 3) // make([]int - array of int, 0 - start_len, 3 - capacity)
	for i := 0; i < cap(slice); i++ {
		slice = slice[0:len(slice)+1]
		slice[i] = i
	}
	for i := range slice {
		fmt.Println("reSlicing element:", slice[i])
	}
}

func	copyAppend() {
	slice_ := []int{1, 2, 3}
	slice := make([]int, 10)
	copy(slice, slice_)
	fmt.Println("Slice_ copy inside slice:", slice)
	slice = append(slice[0:3], 4, 5, 6)
	fmt.Println("Append elements to slice:", slice)
}

func	enlargedSlice(slice []int, factor int) (slice_ []int) {
	slice_ = make([]int, len(slice)*factor)
	copy(slice_, slice)

	return
}

func	insertSlice(slice, insertSlice []string, index int) (slice_ []string) { // my own solution
	slice_ = make([]string, len(slice) + len(insertSlice))
	copy(slice_, slice[:index])
	copy(slice_[index:], insertSlice[:len(insertSlice)])
	copy(slice_[index + len(insertSlice):], slice[index:])
	return
}

func	insertSlice_(slice, insertion []string, index int) []string { // course solution
	result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

type filter_ func(int) bool
func	isEven_(value int) (result bool) {
	if value % 2 == 0 {
		result = true
	}
	return
}
func	filterFunc(slice []int, function filter_) (slice_ []int){
	for _, v := range slice {
		if function(v) {
			slice_ = append(slice_, v)
		}
	}
	return
}

// simulating operations with append block

func	appendOperations() {
	slice := []int{1, 2, 3}
	slice_ := []int{4, 5, 6}

	i := 0
	j := 5
	x := 123

	var popStack, pushStack int

	slice = append(slice, slice_...) // append slice_ to slice
	slice = append(slice[:i], slice[i + 1:]...) // delete item at index i
	slice = append(slice[:i], slice[j:]...) // cut from index i till j out of slice
	slice = append(slice, make([]int, j)...) // extend slice with a new slice_ with size j
	slice = append(slice[:i], append([]int{x}, slice[i:]...)...) // insert x at index i
	slice = append(slice[:i], append(make([]int, j), slice[i:]...)...) // insert new slice of length j at index i
	slice = append(slice[:i], append(slice_, slice[i:]...)...) // insert and existing slice_ at index I in slice

	popStack, slice = slice[len(slice) - 1], slice[:len(slice) - 1] // pop highest element from stack (slice)
	pushStack = 345
	slice = append(slice, pushStack) // push an element on stack (slice)
	fmt.Println("Slice operations:", slice, "popStack:", popStack)
}

func	bubbleSort(slice []int) []int{ // my own version
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice) - i - 1; j++ {
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
	return slice
}


func	bubbleSort__(sl []int) { // course version
	// passes through the slice:
	for pass:=1; pass < len(sl); pass++ {
		// one pass:
		for i:=0; i < len(sl) - pass; i++ {		// the bigger value 'bubbles up' to the last position
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
	return
}

func	reverseString(str string) string {
	str_ := []byte(str)
	for i, j := 0, len(str_) - 1; i < j; i, j = i + 1, j - 1 {
		str_[i], str_[j] = str_[j], str_[i]
	}
	return string(str_)
}

func	main() {
	defaultArray()
	fmt.Println("Iterative fibonacci:", iterativeFibonacci())
	defaultSlice()

	sliceFunc_ := []int{1, 2, 3, 4, 5}
	sliceFunc(sliceFunc_) // it changes default values of sliceFunc_
	fmt.Println(sliceFunc_)
	makeSlice()
	fmt.Println(fibSlice(5))
	bufferBytes()
	reSlicing()
	copyAppend()
	enlargedSlice(sliceFunc_, 5)

	strSlice := []string{"M", "N", "O", "P", "Q", "R"}
	strInsertSlice := []string{"A", "B", "C"}
	insertSlice(strSlice, strInsertSlice, 3)
	fmt.Println("Filter:", filterFunc(sliceFunc_, isEven_))
	appendOperations()
	bubbleSort_ := []int{6, 2, 1, 4, 3}
	fmt.Println(bubbleSort(bubbleSort_))
	bubbleSort__(bubbleSort_)
	fmt.Println(reverseString("Hello"))
}