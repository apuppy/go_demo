package main

import (
	"log"
	"time"
)

func main() {
	outerChan := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			outerChan <- 1
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 8; i++ {
			outerChan <- 1
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	{
		counter := 0
		for counter < 16 {
			counter += <-outerChan
			log.Println("counter: ", counter)
		}
	}
}
