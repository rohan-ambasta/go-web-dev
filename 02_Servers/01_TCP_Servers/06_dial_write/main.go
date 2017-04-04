package main

import (
	"fmt"
	"log"
	"net"
)

// To run this program
// Run 02_read first. It will act as server for this example
// Then run this program
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you")
}
