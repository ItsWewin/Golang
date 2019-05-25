package main

import (
	"fmt"
	"log"
	"net"
)

func accept(c net.Conn) {
	fmt.Printf("client: %v connected\n", c.RemoteAddr().String())
	for {
		reader := make([]byte, 1024)
		n, err := c.Read(reader)
		if err != nil {
			fmt.Printf("reader data error\n")
			break
		}
		fmt.Printf(string(reader[:n]))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("addr: %v\n", listener.Addr().String())
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go accept(conn)
	}
}
