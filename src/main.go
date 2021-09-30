package main

import (
	"github.com/apuppy/go_demo/rabbit"
)

func main() {
	rabbit.EmitLogDirect()
	rabbit.ReceiveLogsDirect()
}
