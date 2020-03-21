package main

import (
	"fmt"
	"net/http"
	"urlshort/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{"/godev": "https://go.dev", "/vuejs": "https://vuejs.org"}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /hacker-news
  url: https://news.ycombinator.com/news
- path: /github-trending
  url: https://github.com/trending
`
	yamlHandler, err := handler.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
