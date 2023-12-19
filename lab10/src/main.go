package main

import (
	"log"
	"math"
	"time"
)

const (
	NUMBER_OF_THREADS int = 5
	LEFT              int = 1_000_000
	RIGHT             int = 2_000_000
)

func isPrime(number int) bool {
	for i := 2; i < number-1; i++ {
		if i*i > number {
			return true
		} else if number%i == 0 {
			return false
		}
	}

	return true
}

func betterIsPrime(number int, smallPrimes []int) bool {
	for _, prime := range smallPrimes {
		if number%prime == 0 {
			return false
		} else if prime*prime > number {
			return true
		}
	}

	return true
}

func multithreadedBetterIsPrime(number int, smallPrimes []int) bool {
	ch := make(chan bool, 1)

	for i := 0; i < NUMBER_OF_THREADS; i++ {
		from := i * len(smallPrimes) / NUMBER_OF_THREADS
		to := (i + 1) * len(smallPrimes) / NUMBER_OF_THREADS

		slice := smallPrimes[from:to]
		go func() {
			for _, prime := range slice {
				if number%prime == 0 {
					ch <- false
				} else if prime*prime > number {
					ch <- true
				}
			}
		}()
	}

	return <-ch
}

func calculate(left int, right int, cb func(int, []int) bool) {
	smallPrimes := []int{}
	middle := int(math.Ceil(math.Sqrt(float64(right))))
	for i := 2; i <= middle; i++ {
		if isPrime(i) {
			smallPrimes = append(smallPrimes, i)
		}
	}

	primes := []int{}
	for i := left; i <= right; i++ {
		if cb(i, smallPrimes) {
			primes = append(primes, i)
		}
	}
}

func main() {
	start := time.Now()
	calculate(LEFT, RIGHT, betterIsPrime)

	log.Printf("Sequential: %d", time.Since(start))

	start = time.Now()
	calculate(LEFT, RIGHT, multithreadedBetterIsPrime)

	log.Printf("Sequential: %d", time.Since(start))
}
