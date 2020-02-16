package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func main() {
	// create watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()
	fmt.Println("watcher initialized!")

	//
	if err := filepath.Walk("/Users/hongde/Desktop/test", watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! $#v\n", event)
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
