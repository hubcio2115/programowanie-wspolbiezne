package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
)

const QUEUE_PATH = "./server-queue"

func main() {
	id := flag.String("userId", "12345", "Id of the user you want to get.")
	outFilePath := flag.String("outFifo", "./out", "Path to the fifo file you want to read the answer from")

	flag.Parse()

	os.Remove(*outFilePath)
	err := syscall.Mkfifo(*outFilePath, 0655)
	if err != nil {
		log.Fatal("Couldn't create out queue.")
	}

	serverQueue, err := os.OpenFile(QUEUE_PATH, os.O_WRONLY, 0655)
	if err != nil {
		log.Fatal("Couldn't open server queue.")
	}

	command := fmt.Sprintf("%s %s", *id, *outFilePath)
	_, err = serverQueue.WriteString(command)
	if err != nil {
		log.Fatalf("Couldn't write to queue: %s", err.Error())
	}

	serverQueue.Close()

	outFile, err := os.Open(*outFilePath)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(outFile)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	outFile.Close()
}
