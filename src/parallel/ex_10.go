package main

import (
	"log"
	"sync"
	"time"
)

var Tasks []string
var mutex = &sync.Mutex{}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		Tasks = append(Tasks, "T1")
		mutex.Unlock()
		log.Println(Tasks)
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		Tasks = append(Tasks, "T2")
		mutex.Unlock()
		log.Println(Tasks)
		wg.Done()
	}()

	wg.Wait()
}
