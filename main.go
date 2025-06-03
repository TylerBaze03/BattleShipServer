package main

import (
	"fmt"
	"net"
	"os"
)

const (
	serverAddress = "localhost:8080"
)

func handleConnection(con net.Conn) {
	defer con.Close()
}

func main() {
	listener, err := net.Listen("tcp", serverAddress)

	//check if server is running w/out error
	if err != nil {
		fmt.Println("Error with connection: ", err)
		os.Exit(1)
	}
	fmt.Println("listening at ", serverAddress)

	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("Couldnt connect ", err)
			continue
		}
		go handleConnection(con)
	}

}
