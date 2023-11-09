package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const QUEUE_PATH = "./server-queue"

func readData(filePath *string) map[string]string {
	database, err := os.Open(*filePath)
	if err != nil {
		log.Fatal("Couldn't read the database!")
	}
	defer database.Close()

	res := make(map[string]string)
	scanner := bufio.NewScanner(database)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		res[line[0]] = line[1]
	}

	return res
}

func catchSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGUSR1)

	for {
		sig := <-ch

		if sig == syscall.SIGUSR1 {
			log.Printf("Program terminated with a signal: %v\n", sig)
			os.Exit(0)
		}

		log.Printf("Ignoring signal: %v\n", sig)
	}
}

func main() {
	go catchSignals()

	log.Println(os.Getpid())

	databasePath := flag.String("databasePath", "./database", "Path to the database text file.")

	flag.Parse()

	data := readData(databasePath)

	os.Remove(QUEUE_PATH)
	err := syscall.Mkfifo(QUEUE_PATH, 0655)
	if err != nil {
		log.Fatal("Couldn't create fifo queue!")
	}

	for {
		log.Println("New loop")

		queue, err := os.Open(QUEUE_PATH)
		if err != nil {
			log.Fatal("Couldn't open fifo queue!")
		}

		scanner := bufio.NewScanner(queue)

		for scanner.Scan() {
			command := scanner.Text()

			pair := strings.Split(command, " ")

			name := data[pair[0]]
			path := pair[1]

			outFifo, err := os.OpenFile(path, os.O_WRONLY, 0655)
			if err != nil {
				log.Printf("Couldn't read out queue for: %s", pair[1])
			}

			if name == "" {
				outFifo.WriteString("Nie ma")
			} else {
				outFifo.WriteString(name)
			}

			outFifo.Close()
		}

		queue.Close()
		log.Print("End loop\n\n")
	}
}
