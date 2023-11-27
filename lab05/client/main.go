package main

import (
	"fmt"
	"log"

	"github.com/hslam/ipc"
)

const (
	SERVER_KEY = 3321
	CLIENT_KEY = 3321
)

func main() {
	server_msgid, err := ipc.Msgget(SERVER_KEY, 0666)
	if err != nil {
		log.Fatal("Server is not running.")
	}

	fmt.Print("Input word you want to translate: ")
	var word_to_translate string
	fmt.Scanln(&word_to_translate)

	err = ipc.Msgsend(server_msgid, 1, []byte(word_to_translate), 0666)
	if err != nil {
		panic(err)
	}

	client_msgid, err := ipc.Msgget(CLIENT_KEY, ipc.IPC_CREAT|0666)
	bytes, err := ipc.Msgreceive(client_msgid, 0, 0666)
	if err != nil {
		log.Fatal(err)
	}

	message := string(bytes)

	fmt.Printf("%s -> %s", word_to_translate, message)
}
