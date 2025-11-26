package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %v", err)
			break
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("Error: %v", err)
			break
		}
	}

}

// command:
// "nc -u -l 42069"	--	run this to see the writes as logs on a seperate terminal
