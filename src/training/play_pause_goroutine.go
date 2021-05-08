package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(time.Millisecond * 250)
	i++
	fmt.Println(i)
}

func routine(command chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(time.Second * 1)
	command <- "Pause"

	time.Sleep(time.Second * 1)
	command <- "Play"

	time.Sleep(time.Second * 1)
	command <- "Stop"

	wg.Wait()
}
