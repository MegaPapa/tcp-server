package main

import (
	"net"
	"fmt"
)

func main() {
	serverAddress := SERVER_HOST + ":" + SERVER_PORT
	connection, err := net.Dial(SERVER_PROTOCOL, serverAddress)
	defer connection.Close()
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return;
	}
	fmt.Fprintf(connection, "Asdfasdf")
}
