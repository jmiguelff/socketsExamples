package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide port number")
		return
	}
	portString := ":" + arguments[1]

	// Resolve UDP address -> build IP:PORT pair
	s, err := net.ResolveUDPAddr("udp4", portString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Debug
	fmt.Print("Listen to: ")
	fmt.Println(s)

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Packet from:"+"%s"+":"+"%d\n", addr.IP.String(), addr.Port)
		fmt.Printf("%s", hex.Dump(buf[:n]))
	}
}
