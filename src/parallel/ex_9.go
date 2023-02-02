package main

import (
	"log"
	"sync"
	"time"
)

var Tasks []string

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(1 * time.Second)
		Tasks = append(Tasks, "T1")
		log.Println(Tasks)
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		Tasks = append(Tasks, "T2")
		log.Println(Tasks)
		wg.Done()
	}()

	wg.Wait()
}
