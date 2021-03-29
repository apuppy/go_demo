package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder

	for i := 0; i < 1000; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())
}
