// package main

// import "fmt"

// func square(c chan int) {
// 	fmt.Println("[square] reading")
// 	num := <-c
// 	c <- num * num
// }

// func cube(c chan int) {
// 	fmt.Println("[cube] reading")
// 	num := <-c
// 	c <- num * num * num
// }

// func main() {
// 	fmt.Println("[main] main() started")

// 	squareChan := make(chan int)
// 	cubeChan := make(chan int)

// 	go square(squareChan)
// 	go cube(cubeChan)

// 	testNum := 3
// 	fmt.Println("[main] sent testNum to squareChan")

// 	squareChan <- testNum

// 	fmt.Println("[main] resuming")
// 	fmt.Println("[main] sent testNum to cubeChan")

// 	cubeChan <- testNum

// 	fmt.Println("[main] resuming")
// 	fmt.Println("[main] reading from channels")

// 	squareVal, cubeVal := <-squareChan, <-cubeChan
// 	sum := squareVal + cubeVal

// 	fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
// 	fmt.Println("[main] main() stopped")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// var start time.Time

// func init() {
// 	fmt.Println("init func")
// 	start = time.Now()
// }

// func service1(c chan string) {
// 	time.Sleep(3 * time.Second)
// 	c <- "Hello from service 1"
// }

// func service2(c chan string) {
// 	time.Sleep(5 * time.Second)
// 	c <- "Hello from service 2"
// }

// func main() {
// 	fmt.Println("main() started", time.Since(start))

// 	chan1 := make(chan string)
// 	chan2 := make(chan string)

// 	go service1(chan1)
// 	go service2(chan2)

// 	select {
// 	case res := <-chan1:
// 		fmt.Println("Response from service 1", res, time.Since(start))
// 	case res := <-chan2:
// 		fmt.Println("Response from service 2", res, time.Since(start))
// 	}

// 	fmt.Println("main() stopped", time.Since(start))
// }








package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
