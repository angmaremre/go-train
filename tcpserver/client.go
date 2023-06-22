package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	connection, connectionError := net.Dial("tcp", "localhost:9090")
	if connectionError != nil {
		panic(connectionError.Error())
	}

	defer connection.Close()

	stream, streamErro := io.ReadAll(connection)
	if streamErro != nil {
		fmt.Println(streamErro.Error())
	} else {
		fmt.Println(string(stream))
	}
}
