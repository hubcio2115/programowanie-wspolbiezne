package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hslam/ipc"
)


const (
	SERVER_KEY = 3321
	CLIENT_KEY = 1233
  MAX_MESSAGE_SIZE = 256
)

type Message struct {
  Type uint;
  Text [MAX_MESSAGE_SIZE]byte;
}

func main() {
  pid := uint(os.Getpid())
	server_msgid, err := ipc.Msgget(SERVER_KEY, 0666)
	if err != nil {
		log.Fatal("Server is not running.")
	}

	fmt.Print("Input word you want to translate: ")
	var word_to_translate string
	fmt.Scanln(&word_to_translate)

  err = ipc.Msgsend(server_msgid, pid, []byte(word_to_translate), 0600)
	if err != nil {
		panic(err)
	}

	client_msgid, err := ipc.Msgget(CLIENT_KEY, 0600)
	bytes, err := ipc.Msgreceive(client_msgid, pid, 0600)
	if err != nil {
		log.Fatal(err)
	}

	reply := string(bytes)

	fmt.Printf("%s -> %s", word_to_translate, reply)
}
