package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverIP   string = "127.0.0.1"
	serverPort int    = 3000
	endMessage string = "end"
	rock       string = "r"
	paper      string = "p"
	scissors   string = "s"
)

func sendResponse(server *net.UDPConn, addr *net.UDPAddr, message string) {
	_, err := server.WriteToUDP([]byte(message), addr)
	if err != nil {
		log.Printf("Couldn't send response %v", err)
	}
}

func main() {
	serverAdress := net.UDPAddr{
		Port: serverPort,
		IP:   net.ParseIP(serverIP),
	}

	server, err := net.ListenUDP("udp", &serverAdress)
	if err != nil {
		log.Fatalf("Couldn't create a server: %s\n", err.Error())
	}

	log.Printf("Server listening at %s\n", fmt.Sprintf("%s:%d", serverIP, serverPort))

	var player1 *net.UDPAddr
	var player2 *net.UDPAddr

	buffer := make([]byte, 64)
	for {
		// Wait for first package of each player to start the game
		_, player1, err = server.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Couldn't read from UDP: %s\n", err.Error())
			continue
		}
		log.Println("Player1 connected")

		_, player2, err = server.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Couldn't read from UDP: %s\n", err.Error())
			player1 = nil
			continue
		}
		log.Println("Player2 connected")

		log.Println("Game started")
	game:
		for {
			var playerOneMove string
			var playerTwoMove string
			var playerOneScore int
			var playerTwoScore int

			log.Println("Waiting for moves...")
			for range [2]int{} {
				n, adress, err := server.ReadFromUDP(buffer)
				if err != nil {
					log.Printf("Something went wrong: %s\n", err.Error())
				}

				message := string(buffer[:n])

				switch adress.String() {
				case player1.String():
					playerOneMove = message
				case player2.String():
					playerTwoMove = message
				}
			}

			// If one of the players sent "end" send to the other one that game ended
			if playerOneMove == endMessage {
				sendResponse(server, player2, endMessage)
				break game
			}

			if playerTwoMove == endMessage {
				sendResponse(server, player1, endMessage)
				break game
			}

			// Calculate what happened in a turn and send points
			turn := playerOneMove + playerTwoMove
			if turn == "rr" || turn == "ss" || turn == "pp" {
				sendResponse(server, player1, "0")
				sendResponse(server, player2, "0")

				log.Println("Draw")
			} else if turn == "rs" || turn == "sp" || turn == "pr" {
				sendResponse(server, player1, "1")
				sendResponse(server, player2, "0")

				playerOneScore++
				log.Println("Player1 scores a point")
			} else if turn == "rp" || turn == "sr" || turn == "ps" {
				sendResponse(server, player1, "0")
				sendResponse(server, player2, "1")

				playerTwoScore++
				log.Println("Player2 scores a point")
			}

			log.Printf("Player1: %d, Player2: %d\n", playerOneScore, playerTwoScore)
		}

		player1 = nil
		player2 = nil
		log.Println("Game Ended")
	}
}
