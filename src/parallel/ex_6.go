package main

import (
	"github.com/remeh/sizedwaitgroup"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	counter := 0
	var swg = sizedwaitgroup.New(6)

	for i := 1; i <= 16; i++ {
		swg.Add()
		counter++
		go func(thisCounter int) {
			log.Println("Start", thisCounter)
			swg.Done()
		}(counter)
	}

	swg.Wait()

	elapsed := time.Since(start)
	log.Println("elapsed time: ", elapsed)
}
