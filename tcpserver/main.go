package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, listenerError := net.Listen("tcp", ":9090")

	if listenerError != nil {
		panic(listenerError.Error())
	}
	defer listener.Close()

	for {
		connection, connectionError := listener.Accept()
		if connectionError != nil {
			fmt.Println(connectionError.Error())
			connection.Close()
			continue
		} else {
			io.WriteString(connection, "Merhaba nbr")
			connection.Close()
		}
	}
}
