package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverIP   = "127.0.0.1"
	serverPort = 3000
	endMessage = "end"
	rock       = "r"
	paper      = "p"
	scissors   = "s"
)

func main() {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		log.Fatalf("Couldn't create connection: %v", err)
		return
	}
	defer conn.Close()

	score := 0
	buffer := make([]byte, 64)

	// Send a message to start a game
	conn.Write([]byte(""))
game:
	for {
		var move string
		fmt.Print("Input your move (r = rock, p = paper, s = scissors, end = end game): ")
		fmt.Scanln(&move)

		if move == endMessage {
			conn.Write([]byte(endMessage))
			fmt.Printf("Your final score: %d\n", score)
			break game
		}

		isMoveValid := move != rock || move != paper || move != scissors
		if !isMoveValid {
			fmt.Println("Invalid move!")
			continue
		}

		conn.Write([]byte(move))

		n, _ := conn.Read(buffer)

		response := string(buffer[:n])
		switch response {
		case "1":
			score++
		case endMessage:
			fmt.Printf("Your final score: %d\n", score)
			break game
		}

		fmt.Printf("Your score: %d\n", score)
	}
}
