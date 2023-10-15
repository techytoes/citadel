package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read the query from the client
	queryBuffer := make([]byte, 1024)
	n, err := conn.Read(queryBuffer)
	if err != nil {
		// Handle the error
		return
	}

	query := string(queryBuffer[:n])

	// Process the query (e.g., execute it against your database)
	// Simulate a response for demonstration purposes
	response := "Query result for: " + query

	// Logging query and response for easier debugging
	fmt.Printf("Query: %s\nResponse: %s\n", query, response)
	_, err = conn.Write([]byte(response))
	if err != nil {
		// Handling error: JUST PANIC FOR NOW!!
		panic(err)
	}
}

func main() {
	socketPath := "/tmp/db_server_new.sock"

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// Handle termination signals (e.g., Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			_ = l.Close()         // Close the socket
			os.Remove(socketPath) // Remove the socket file
			_ = sig               // Handle the signal
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleClient(conn)
	}
}
