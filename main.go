package main

import (
	"fmt"
	"net"
	"os"
	"battleshipServer/src"
)

const (
	serverAddress = "localhost:8080"

	DEBUG = true
)


func makePlayer() *battleShip.Player{
	newPlay := battleShip.NewPlayer("Tyler")

	return newPlay
}



func handleConnection(con net.Conn) {
	defer con.Close()
	/*
	if (DEBUG){
		newPlay := makePlayer()
		gameInfo := battleShip.NewGame(newPlay)
		// Need to create a lobby system so players could join and wait for player to join their game
	}
	*/
	
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
