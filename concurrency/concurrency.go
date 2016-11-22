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

	waitGrp.Wait()
}
