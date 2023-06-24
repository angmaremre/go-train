package main

import (
	"bufio"
	"fmt"
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
			scanner := bufio.NewScanner(connection)
			for scanner.Scan() {
				var line = scanner.Text()
				if line == "" {
					break
				}
				fmt.Println(line)
			}
			//io.WriteString(connection, "Merhaba nbr")
			connection.Close()
		}
	}
}
