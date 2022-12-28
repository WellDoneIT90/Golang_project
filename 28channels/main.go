package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)
	wg.Add(2)
	// receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-myChannel)

		wg.Done()
	}(myChannel, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 0
		// myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
