package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var localAddr *string = flag.String("l", "localhost:9999", "local address")
var remoteAddr *string = flag.String("r", "localhost:80", "remote address")

func main() {
	flag.Parse()

	fmt.Printf("Listening: %v\n Proxying: %v\n\n", *localAddr, *remoteAddr)

	listener, err := net.Listen("tcp", *localAddr)
	if err != nil {
		panic(err)
	}
	for {
		connProxy, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}

		go func() {
			connRemote, err := net.Dial("tcp", *remoteAddr)
			if err != nil {
				log.Println("error dailing remote addr", err)
				return
			}
			go io.Copy(connRemote, connProxy)
			io.Copy(connProxy, connRemote)
			connRemote.Close()
			connProxy.Close()
		}()
	}
}
