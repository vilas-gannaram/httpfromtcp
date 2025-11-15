package main

import (
	"bytes"
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

	str := ""

	for {
		data := make([]byte, 8)

		n, err := file.Read(data)
		if err != nil {
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]

			fmt.Println(str)
			str = ""
		}

		str += string(data)

	}

	if str != "" {
		fmt.Println(str)
	}
}
