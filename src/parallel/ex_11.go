package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// !!! Caution: maybe this program is a bad case! The style to use for loop.

var Tasks []string
var mutex sync.Mutex
var lastDump int64 = 0

func main() {
	for i := 0; i < 32; i++ {
		go func(iteration int) {
			duration := rand.Intn(5000)
			time.Sleep(time.Duration(duration) * time.Millisecond)

			Add(iteration)
		}(i)
	}

	// for死循环，这种写法要不得，测试过程中发现CPU被打满
	for {
	}
}

func Add(item int) {
	mutex.Lock()
	Tasks = append(Tasks, fmt.Sprintf("T%d", item))
	if len(Tasks) == 10 {
		taskToDump := Tasks[:10]
		Tasks = Tasks[10:]
		Dump(taskToDump)
	} else {
		now := time.Now().Unix()
		go func() {
			time.Sleep(5 * time.Second)
			mutex.Lock()
			if now-lastDump > 5 && len(Tasks) > 0 {
				lastDump = time.Now().Unix()
				tasksToDump := Tasks
				Tasks = []string{}
				Dump(tasksToDump)
			}
			mutex.Unlock()
		}()
	}
	mutex.Unlock()
}

func Dump(tasksToDump []string) {
	log.Println("Dumping: ", tasksToDump)
}
