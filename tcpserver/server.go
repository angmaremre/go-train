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
			var responseBody = "<html><title></title><body>FLOOOO</body></html>"
			connection.Write([]byte("HTTP/1.1 200 OK\r\n"))
			connection.Write([]byte(fmt.Sprintf("Contect-Lenght: #{len(responseBody)}\r\n")))
			fmt.Fprintf(connection, "Content-Type: text/html\r\n")
			fmt.Fprintf(connection, "\r\n")
			fmt.Fprintf(connection, responseBody)
			fmt.Println("")
			connection.Close()
		}
	}
}
