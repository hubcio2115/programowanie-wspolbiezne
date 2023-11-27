package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unsafe"

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
	file, err := os.OpenFile("./dictionary.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	dictionary := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		dictionary[line[0]] = line[1]
	}

	log.Println("Server started")

	server_msgid, err := ipc.Msgget(SERVER_KEY, ipc.IPC_CREAT|ipc.IPC_EXCL|0666)
  if err != nil {
    log.Panic(err)
  }
	defer ipc.Msgrm(server_msgid)


	client_msgid, err := ipc.Msgget(CLIENT_KEY, ipc.IPC_CREAT|0666)
  if err != nil {
    log.Panic(err)
  }
	defer ipc.Msgrm(client_msgid)

  message:= Message{}

	for {
		length, err := ipc.Msgrcv(server_msgid, uintptr(unsafe.Pointer(&message)), MAX_MESSAGE_SIZE, 0, 0600)
		if err != nil {
			log.Panic(err)
		}

    fmt.Printf("Received: %s\n", string(message.Text[:length]))

    value := dictionary[string(message.Text[:length])]
		if value == "" {
			value = "Nie ma w słowniku."
		} 


		err = ipc.Msgsend(client_msgid, message.Type, []byte(value), 0600)
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(2 * time.Second)
	}
}
