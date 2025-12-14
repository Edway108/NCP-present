package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// TCP Echo Server - receives messages and sends them back
func startTCPServer() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on :8080")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle connection in goroutine
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Client connected from %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received: %s\n", message)

		// Echo back the message in uppercase
		response := strings.ToUpper(message) + "\n"
		conn.Write([]byte(response))
	}

	fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
}

func main() {
	startTCPServer()
}
