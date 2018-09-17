package main

import (
	"net"
	"fmt"
	"tcp-server/server"
)

const(
	SERVER_PORT     = "9500"
	SERVER_HOST     = "localhost"
	SERVER_PROTOCOL = "tcp"
)

func main() {
	address := SERVER_HOST + ":" + SERVER_PORT
	listener, err := net.Listen(SERVER_PROTOCOL, address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Printf("Server started at %s:%s\n", SERVER_HOST, SERVER_PORT)
	for {
		connection, err := listener.Accept()
		fmt.Println("New user connected...")
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return
		}
		go server.Handle(connection)
	}
}