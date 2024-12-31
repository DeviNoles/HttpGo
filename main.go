package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello, Go!")
	var uri string = "localhost"
	var port string = "8080"
	l, err := net.Listen("tcp", uri+":"+port)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Printf("Listening on %s:%s\n", uri, port)
}
