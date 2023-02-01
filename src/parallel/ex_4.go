package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	counter := 0

	for i := 1; i <= 16; i++ {
		counter++
		go func(thisCounter int) {
			log.Println("Start", thisCounter)
		}(counter)
	}

	time.Sleep(2 * time.Second)

	elapsed := time.Since(start)
	log.Println("elapsed time: ", elapsed)
}
