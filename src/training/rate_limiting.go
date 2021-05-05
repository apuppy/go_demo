package main

import (
	"fmt"
	"time"
)

func main() {
	// assume that 5 request at a time
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// average limiter, everty 1 second the limiter returns
	limiter := time.Tick(time.Second * 1)

	// handle request by regular rate limiter
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// bursty rate limiter when initialize
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// then every 1 second, burstyLimiter returns(bursty to regular)
	go func() {
		for t := range time.Tick(time.Second * 1) {
			burstyLimiter <- t
		}
	}()

	// assume that 5 requests at a time
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// perform request use a limiter, bursty at begining, then regular
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
