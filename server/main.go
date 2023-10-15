package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Handle database queries and responses here
	// Example: Read query from the client, process it, and send back a response.
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
