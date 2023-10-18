package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	serverBufferPath := flag.String("serverBuffer", "./buffer.txt", "Path to the server buffer file.")
	outFilePath := flag.String("outFile", "./out.txt", "Path to the file with the server answer.")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Lines:")
	lines := []string{*outFilePath}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 1 && line[0] == '\x1D' {
			break
		}

		lines = append(lines, line)
	}

	var serverBuffer *os.File
	var err error
	for {
		serverBuffer, err = os.OpenFile(*serverBufferPath, os.O_EXCL|os.O_WRONLY|os.O_CREATE, 0655)
		if err == nil {
			fmt.Println("Couldn't access server since it's bussy.")
			break
		}
		time.Sleep(time.Millisecond * 200)
	}
	defer serverBuffer.Close()

	for _, line := range lines {
		serverBuffer.WriteString(line + "\n")
	}

	time.Sleep(time.Second)

	outFile, err := os.OpenFile(*outFilePath, os.O_RDONLY, 0655)
	if err != nil {
		log.Panic(err)
	}
	defer outFile.Close()

	reader := bufio.NewScanner(outFile)
	fmt.Println()
	for reader.Scan() {
		fmt.Println(reader.Text())
	}

	os.Remove(*serverBufferPath)
}
