package main

import (
	"net"
	"fmt"
)

func getIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error recovering IP: ", err.Error())
	}
	ipNet, _ := addrs[7].(*net.IPNet)
	return ipNet.IP
}

func main() {
	myIP := getIP()
	fmt.Println(myIP)
}