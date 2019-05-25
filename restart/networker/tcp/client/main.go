package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	address := "127.0.0.1:3000"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Connect %v error: %v", address, err)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')

		newLine := strings.Trim(line, "\r\n")
		if strings.ToLower(newLine) == "exit" {
			fmt.Printf("log out!")
			break
		}

		if err != nil {
			log.Fatal("some error")
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Printf("send data to server error: %v\n", err)
		}
	}
}
