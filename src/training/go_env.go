package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("BABY_NAME", "miao")
	fmt.Println("baby name: ", os.Getenv("BABY_NAME"))
	fmt.Println("foo: ", os.Getenv("FOO"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s =  %s \n", pair[0], pair[1])
	}
}
