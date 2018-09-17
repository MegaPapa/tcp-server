package server

import (
	"net"
	"fmt"
)

const(
	_STANDARD_BUFFER_SIZE = 1024
)

func Handle(connection net.Conn) {
	buffer := make([]byte, _STANDARD_BUFFER_SIZE)
	dataLength, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	data := string(buffer[:dataLength])
	fmt.Printf("Client send next data: %s\n", data)
	defer connection.Close()
}
