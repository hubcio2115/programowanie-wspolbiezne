package main

import (
	"bufio"
	"log"
	"os"
	"syscall"
)

func main() {
	log.Println("Opening queue...")
	queue, _ := os.OpenFile("./server-queue", os.O_WRONLY, 0655)

	log.Println("Opening writer...")
	writer := bufio.NewWriter(queue)
	writer.WriteString("Chuj kurwa")
	writer.Flush()

	log.Println("Opening out...")
	os.Remove("./out")
	_ = syscall.Mkfifo("./out", 0655)
	out, _ := os.Open("./out")
	scanner := bufio.NewScanner(out)

	log.Println("Opening scanner...")
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
