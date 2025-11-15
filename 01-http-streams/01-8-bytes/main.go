package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalf("Failed to Open file %v", err)
	}

	defer file.Close()

	for {
		data := make([]byte, 8)

		n, err := file.Read(data)
		if err != nil {
			break
		}

		fmt.Println(string(data[:n]))

	}
}
