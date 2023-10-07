package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	IN_FILE_PATH  = "./in.txt"
	OUT_FILE_PATH = "./out.txt"
)

func main() {
	inFile, inFileErr := os.OpenFile(IN_FILE_PATH, os.O_WRONLY|os.O_TRUNC, 0644)
	if inFileErr != nil {
		log.Panic("Error while opening/creating in.txt:", inFileErr)
	}
	defer inFile.Close()

	fmt.Print("Enter a value to be calculated: ")

	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal("Couldn't read the value from user.")
	}

	inFile.WriteString(userInput)

	time.Sleep(time.Second)

	outFile, outFileErr := os.OpenFile(OUT_FILE_PATH, os.O_RDONLY, 0644)

	if outFileErr != nil {
		log.Fatal("Error while opening/creating in.txt:", inFileErr)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(outFile)
	if scanner.Scan() {
		line := scanner.Text()

		fmt.Printf("Result: %s", line)
	}
}
