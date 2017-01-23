package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listening...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(msg)
}
