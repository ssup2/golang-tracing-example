package net

import (
	"log"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "7070"
	SERVER_TYPE = "tcp"
)

func Receive() {
	// Listen
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer server.Close()

	for {
		// Accept connection
		conn, _ := server.Accept()

		// Receive message
		go func(conn net.Conn) {
			for {
				buffer := make([]byte, 1024)
				len, err := conn.Read(buffer)
				if err != nil {
					panic(err)
				}
				log.Println("Received:", string(buffer[:len]))
			}
		}(conn)
	}
}

func Send() {
	// Connect to server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send message
	for {
		_, err := conn.Write([]byte("Message"))
		if err != nil {
			panic(err)
		}
	}
}
