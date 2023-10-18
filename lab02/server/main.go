package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	for {
		buffer, err := os.OpenFile("./buffer.txt", os.O_RDONLY, 0655)
		if err != nil {
			time.Sleep(time.Microsecond * 500)
			continue
		}

		stat, _ := buffer.Stat()
		if stat.Size() == 0 {
			buffer.Close()
			time.Sleep(time.Millisecond * 200)
			continue
		}

		scanner := bufio.NewScanner(buffer)

		lines := make([]string, 0)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		outFile, err := os.OpenFile(lines[0], os.O_CREATE|os.O_WRONLY, 0655)
		if err != nil {
			log.Fatal(err)
			buffer.Close()
			continue
		}

		for _, line := range lines[1:] {
			fmt.Println(line)
			outFile.WriteString(line + "\n")
		}
		fmt.Println()

		outFile.Close()
		buffer.Close()

		time.Sleep(time.Second)
	}
}
