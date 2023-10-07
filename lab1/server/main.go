package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	IN_FILE_PATH  = "./in.txt"
	OUT_FILE_PATH = "./out.txt"
)

func doCalculations(num int) int {
	return num*num + 3*num + 6
}

func main() {
	for {
		fmt.Println("Listening...")

		inFile, inFileErr := os.OpenFile(IN_FILE_PATH, os.O_RDONLY|os.O_CREATE, 0644)
		if inFileErr != nil {
			log.Panic("Error while opening/creating in.txt:", inFileErr)
		}
		defer inFile.Close()

		outFile, outFileErr := os.OpenFile(OUT_FILE_PATH, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if outFileErr != nil {
			log.Fatal("Error while opening/creating in.txt:", inFileErr)
		}

		var inputValue int
		scanner := bufio.NewScanner(inFile)
		if scanner.Scan() {
			line := scanner.Text()

			parsedLine, err := strconv.Atoi(line)
			if err != nil {
				outFile.WriteString("Error while parsing the input. It probably cannot be parsed into an integer.")

				inFile.Close()
				outFile.Close()
				continue
			}

			inputValue = doCalculations(parsedLine)
		}

		outFile.WriteString(fmt.Sprintf("%d", inputValue))

		inFile.Close()
		outFile.Close()

		time.Sleep(time.Second)
	}
}
