package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("hahaha!!!")
	ding()
	panic("error defined by myself in main...")
}

func ding() {
	ping()
	panic("error from ding...")
}

func ping() {
	panic("error from ping...")
}
