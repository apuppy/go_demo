package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	counter := 0
	for i := 1; i <= 16 ; i++ {
		go func(){
			log.Println("Start", counter)
			counter++
		}()
	}
	elapsed := time.Since(start)
	log.Println("elapsed time: ", elapsed)
}
