package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in go")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		for val := range myCh {
			fmt.Println(val)
		}
	}(myCh, wg)

	wg.Add(1)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		myCh <- 7
		myCh <- 5
		myCh <- 6
		close(myCh)
	}(myCh, wg)

	wg.Wait()
}
