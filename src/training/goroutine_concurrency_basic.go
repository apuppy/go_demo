package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 10
	guard := make(chan struct{}, maxGoroutines)

	for i := 0; i < 30; i++ {
		wg.Add(1)
		guard <- struct{}{}
		go func(n int) {
			worker(n, &wg)
			<-guard
		}(i)
	}
	// time.Sleep(time.Second * 3)
	wg.Wait()
}

func worker(i int, wg *sync.WaitGroup) {
	fmt.Println("doing work on", i)
	wg.Done()
}
