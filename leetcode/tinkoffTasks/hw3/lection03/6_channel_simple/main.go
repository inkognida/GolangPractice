package main

import (
	"fmt"
	"sync"
)

func main() {
	unbuffered := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		i := 0
		for {
			//time.Sleep(time.Microsecond * 1)
			print("write to\n")

			unbuffered <- fmt.Sprintf("Hello #%d", i)
			i++
		}
		close(unbuffered)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-unbuffered
			if !ok {
				fmt.Println("stop reader")
				//return
			}
			print(v, "\n")
		}
	}()

	wg.Wait()
}
