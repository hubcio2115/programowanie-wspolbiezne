package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number == 2 {
		return true
	}

	if number%2 == 0 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(number)))
	for i := 3; i <= maxDivisor; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func doCalculations(result *chan int, input []int, wg *sync.WaitGroup) {
	for _, value := range input {
		if isPrime(value) {
			*result <- value
		}
	}

	wg.Done()
}

func main() {
	numberOfThreads := flag.Int("threads", 5, "Define how many threads do you want to use to calculate the value.")
	left := flag.Int("left", 5, "Left boundry of the range.")
	right := flag.Int("right", 20, "Right boundry of the range.")
	flag.Parse()

	values := make([]int, 0)
	for value := *left; value < *right; value++ {
		values = append(values, value)
	}

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(*numberOfThreads)
	go func() {
		for i := 0; i < *numberOfThreads; i++ {
			from := i * len(values) / *numberOfThreads
			to := (i + 1) * len(values) / *numberOfThreads

			go doCalculations(&ch, values[from:to], &wg)
		}

		wg.Wait()
		close(ch)
	}()

	result := []int{}
	for value := range ch {
		result = append(result, value)
	}

	fmt.Println(result)
}
