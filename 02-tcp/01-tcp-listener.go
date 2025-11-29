package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

// getLines reads from io.ReadCloser and emits complete lines via a channel
// Lines are delimited by '\n'. Any remaining data after the last newline is also emitted.
func getLines(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		str := ""

		for {
			data := make([]byte, 8) // Small buffer for demonstration (use 1024+ in production)

			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]

			// Check if we have a complete line
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]

				out <- str
				str = ""
			}

			str += string(data)
		}

		// Send any remaining data that didn't end with '\n'
		if str != "" {
			out <- str
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// ========================================================================
	// OPTION 1: Single Client (Server exits after first client disconnects)
	// ========================================================================
	// conn, err := listener.Accept() // Accept() called ONCE
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
	// defer conn.Close()
	//
	// for line := range getLines(conn) {
	// 	fmt.Println(line)
	// }
	// // After client disconnects, main() ends → Server stops
	//
	// Timeline:
	//   Server starts → Client A connects → Handle A → A disconnects → SERVER EXITS
	//   ❌ Client B cannot connect (server already stopped)

	// ========================================================================
	// OPTION 2: Multiple Clients (Sequential - one at a time)
	// ========================================================================
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatalf("Error: %v", err)
	// 	}
	//
	// 	// ❌ BUG: defer in loop doesn't close until main() exits
	// 	// defer conn.Close() // WRONG - causes connection leak
	//
	// 	for line := range getLines(conn) {
	// 		fmt.Println(line)
	// 	}
	//
	// 	conn.Close() // ✓ Close immediately after handling
	// }
	//
	// Timeline:
	//   Client A connects → Handle A → A disconnects
	//                                 ↓
	//                     Client B connects → Handle B → B disconnects
	//                                                   ↓
	//                                       Client C connects → ...
	//
	// Problem: Clients must wait for previous client to finish

	// ========================================================================
	// OPTION 3: Concurrent (Multiple clients simultaneously with goroutines)
	// ========================================================================
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Launch goroutine to handle this connection
		// Pass conn as parameter to avoid variable capture issues
		go func(c net.Conn) {
			defer c.Close() // Each goroutine closes its own connection

			for line := range getLines(c) {
				fmt.Printf("[%s] %s\n", c.RemoteAddr(), line)
			}
		}(conn)
	}
	//
	// Timeline:
	//   Client A connects → goroutine 1 handles A (ongoing...)
	//        ↓
	//   Client B connects → goroutine 2 handles B (ongoing...)
	//        ↓
	//   Client C connects → goroutine 3 handles C (ongoing...)
	//
	// All clients handled simultaneously! Each gets its own goroutine.
}

// ============================================================================
// KEY CONCEPTS
// ============================================================================
//
// 1. GOROUTINES:
//    - Lightweight threads managed by Go runtime
//    - Started with 'go func()'
//    - Enable concurrent execution
//    - Very cheap (can run thousands simultaneously)
//
// 2. WHY PASS conn AS PARAMETER:
//    go func(c net.Conn) { ... }(conn)  // ✓ Correct - each goroutine gets copy
//    go func() { use conn }()            // ❌ Wrong - all share same variable
//
// 3. DEFER IN LOOPS:
//    for { defer f() }  // ❌ Defers accumulate, execute when function returns
//    for { f() }        // ✓ Executes immediately each iteration
//
// 4. PRODUCTION IMPROVEMENTS:
//    - Use larger buffer (1024+ bytes instead of 8)
//    - Add timeout handling
//    - Implement graceful shutdown
//    - Add connection limits to prevent resource exhaustion
//    - Better error logging (don't log.Fatal in goroutines)
// ============================================================================
