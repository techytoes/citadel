package main

import (
	"fmt"
	"net"
)

func main() {
	socketPath := "/tmp/db_server_new.sock"

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send database queries to the server and read responses here
	query := "SELECT * FROM my_table" // Replace with user input

	// Send the query to the server
	_, err = conn.Write([]byte(query))
	if err != nil {
		// Handle the error
	}

	// Receive and display the response from the server
	responseBuffer := make([]byte, 1024)
	n, err := conn.Read(responseBuffer)
	if err != nil {
		// Handle the error
		return
	}

	response := string(responseBuffer[:n])
	fmt.Println("Server response:", response)
}
