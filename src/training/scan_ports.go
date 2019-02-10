package main

import (
	"fmt"
	"net"
	"time"
)

func main(){

	ip := "192.168.31.1"
	var ports = []int{22, 80, 443, 444, 65, 3389}

	fmt.Print("Enter IP address to scan: ")
	fmt.Scan(&ip)

	fmt.Println(ip)

	for _, port := range ports {
		host_ip := fmt.Sprintf("%s:%d", ip, port)

		_, err := net.DialTimeout("tcp", host_ip, 10*time.Millisecond)

		if err != nil {
			fmt.Printf("Port %d is closed \n", port)
		} else {
			fmt.Printf("Port %d is open \n", port)
		}
	}
}