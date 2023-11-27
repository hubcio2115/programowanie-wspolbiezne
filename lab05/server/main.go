package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hslam/ipc"
)

const (
	SERVER_KEY = 3321
	CLIENT_KEY = 3321
)

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

	server_msgid, _ := ipc.Msgget(SERVER_KEY, ipc.IPC_CREAT|ipc.IPC_EXCL|0666)
	defer ipc.Msgrm(server_msgid)
	client_msgid, _ := ipc.Msgget(CLIENT_KEY, ipc.IPC_CREAT|0666)
	defer ipc.Msgrm(client_msgid)

	for {
		bytes, err := ipc.Msgreceive(server_msgid, 1, 0666)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Received: %s\n", string(bytes))

		value := dictionary[string(bytes)]
		var message string
		if value == "" {
			message = "Nie ma w słowniku."
		} else {
			message = value
		}

		err = ipc.Msgsend(client_msgid, 1, []byte(message), 0666)
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(2 * time.Second)
	}
}
