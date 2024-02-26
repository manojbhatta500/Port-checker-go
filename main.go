package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var rules = []string{
		"1. you can put Ip addresses of any device eg. Mobile,Laptop,Router etc.\n(any device which has ip address)",
		"2. you can test Websites also here simply copy and paste the url.",
		"3. it will scan the all the ports and returs the opened ports at the end:",
	}
	_ = rules

	fmt.Println("hello welcome to manoj port checker app")

	scanAllPorts("192.168.1.85")

	// "192.168.0.12"  my ip

}

func scanningport() {
	fmt.Println("starting the scanning")

	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connection was successfull")
	}
}

// 1 to 65535  are total ports

func scanAllPorts(host string) {
	var openPorts []int

	timeout := time.Millisecond * 500 // Timeout for port scan

	for port := 1; port <= 65535; port++ { // Change the range as needed
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			fmt.Printf("Port %d is closed\n", port)
		} else {
			conn.Close()
			openPorts = append(openPorts, port) // Append open port to openPorts slice
			fmt.Printf("Port %d is open\n", port)
		}
	}

	fmt.Println("Open ports:", openPorts)
}
