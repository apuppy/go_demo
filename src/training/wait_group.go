package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d starting \n", i)
	time.Sleep(time.Second * 1)
	fmt.Printf("worker %d done \n", i)
}
