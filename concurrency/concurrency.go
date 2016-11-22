package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//Parallelism
	runtime.GOMAXPROCS(2)

	//Concurrency
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Ray")
	}()

	//Wait that all goroutines in the waitGrp are done
	waitGrp.Wait()

	//Channels
	ch := make(chan string)
	doneCh := make(chan bool)

	go abcGen(ch)
	go printer(ch, doneCh)

	//Exit the main function when the goroutine "printer" is done
	<-doneCh
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= ('z'); l++ {
		ch <- string(l)
	}

	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}

	doneCh <- true
}
