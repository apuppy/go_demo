package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	beginTime := time.Now()


	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go workerWrong(i, wg)
		go workerRight(i, &wg)
	}

	wg.Wait()

	elapsed := time.Since(beginTime)
	log.Println("elapsed time: ", elapsed)
}

// TODO explain the behavior when understand the detail under the hood
func workerWrong(i int, wg sync.WaitGroup) {
	fmt.Printf(">>> counter: %d \n", i)
	wg.Done()
}

func workerRight(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	fmt.Printf(">>> counter: %d \n", i)
	wg.Done()
}
