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
		log.Fatalf("Error resolving address: %v", err)
	}

	dial, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error creating UDP connection: %v", err)
	}
	defer dial.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("UDP client started. Type messages to send:")

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		n, err := dial.Write([]byte(line))
		if err != nil {
			log.Printf("Error sending: %v", err)
			continue // Try to send next message
		}

		fmt.Printf("Sent %d bytes\n", n)
	}
}

// Start a UDP server (listener/receiver): "nc -u -l 42069"
