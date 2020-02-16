package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// create file wather
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()
	fmt.Println("file watcher initialized!")

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVEN! %#v\n", event)
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Printf("ERROR", err)
			}
		}
	}()

	// watch a file or directory
	if err := watcher.Add("/Users/hongde/code/work/crawler_api/runtime/logs/app.log"); err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Println("hey")

	<-done
}
