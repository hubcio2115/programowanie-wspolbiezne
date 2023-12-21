package main

import (
	"log"
	"math"
	"sync"
	"time"
)

const (
	NUMBER_OF_THREADS int = 5
	LEFT              int = 1_000_000
	RIGHT             int = 2_000_000
	CHUNK_SIZE        int = (RIGHT - LEFT) / NUMBER_OF_THREADS
)

func isPrimePart1(number int) bool {
	for i := 2; i < number-1; i++ {
		if i*i > number {
			return true
		} else if number%i == 0 {
			return false
		}
	}

	return true
}

func isPrimePart2(number int, smallPrimes []int) bool {
	for _, prime := range smallPrimes {
		if number%prime == 0 {
			return false
		} else if prime*prime > number {
			return true
		}
	}

	return true
}

func calculate(left int, right int) []int {
	smallPrimes := []int{}
	middle := int(math.Ceil(math.Sqrt(float64(right))))
	for i := 2; i <= middle; i++ {
		if isPrimePart1(i) {
			smallPrimes = append(smallPrimes, i)
		}
	}

	primes := []int{}
	for i := left; i <= right; i++ {
		if isPrimePart2(i, smallPrimes) {
			primes = append(primes, i)
		}
	}

	return primes
}

func multithreadedCalculate(left int, right int) {
	smallPrimes := []int{}
	middle := int(math.Ceil(math.Sqrt(float64(right))))
	for i := 2; i <= middle; i++ {
		if isPrimePart1(i) {
			smallPrimes = append(smallPrimes, i)
		}
	}

	primes := []int{}
	wg := sync.WaitGroup{}
	wg.Add(NUMBER_OF_THREADS)
	mx := sync.Mutex{}

	for i := 0; i < NUMBER_OF_THREADS; i++ {
		value := i
		go func() {
			for j := LEFT + value*CHUNK_SIZE; j <= LEFT+(value+1)*CHUNK_SIZE; j++ {
				if isPrimePart2(j, smallPrimes) {
					mx.Lock()
					primes = append(primes, j)
					mx.Unlock()
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	start := time.Now()
	calculate(LEFT, RIGHT)
	end := time.Since(start)

	log.Printf("Sequential: %dms", end)

	start = time.Now()
	multithreadedCalculate(LEFT, RIGHT)
	end = time.Since(start)

	log.Printf("Parallel:   %dms", end)
}
