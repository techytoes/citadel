package main

import (
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
}
