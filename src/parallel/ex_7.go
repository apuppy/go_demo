package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 8; i++ {
			counter++
			time.Sleep(100 * time.Millisecond)
			log.Println("Thread 1: ", counter)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 8; i++ {
			counter++
			time.Sleep(100 * time.Millisecond)
			log.Println("Thread 2: ", counter)
		}
		wg.Done()
	}()

	wg.Wait()
}
