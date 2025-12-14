package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server!")
	fmt.Println("Type messages (type 'quit' to exit):")

	// Read from stdin and send to server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "quit" {
			break
		}

		// Send message to server
		fmt.Fprintf(conn, "%s\n", message)

		// Read response
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			break
		}

		fmt.Printf("Server response: %s", response)
	}
}
