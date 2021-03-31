package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at: ", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stoppped")
}
