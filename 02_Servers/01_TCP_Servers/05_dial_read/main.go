package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// To run this program
// Run 01_write first. It will act as server for this example
// Then run this program
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
