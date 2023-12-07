package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func doCalculations(result *int, input []int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	partialResult := 0
	for _, value := range input {
		partialResult += value
	}

	mutex.Lock()
	*result += partialResult
	mutex.Unlock()

	wg.Done()
}

func main() {
	numberOfThreads := flag.Int("threads", 5, "Define how many threads do you want to use to calculate the value.")
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal("You must provide one input file.")
	}

	file, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalf("Couldn't open specified file: %s", err.Error())
	}

	values := make([]int, 0)
	for _, line := range file {
		value, _ := strconv.Atoi(strings.TrimRight(string(line), "\n"))

		values = append(values, value)
	}

	result := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(*numberOfThreads)
	for i := 0; i < *numberOfThreads; i++ {
		from := i * len(values) / *numberOfThreads
		to := (i + 1) * len(values) / *numberOfThreads

		go doCalculations(&result, values[from:to], &mutex, &wg)
	}

	wg.Wait()
	fmt.Println(result)
}
