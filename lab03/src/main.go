package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

const regexPattern string = `\\input{([^}]+)\}`

func readFromFile(path string, regex *regexp.Regexp, pattern *string, c chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	file, err := os.Open(regex.FindStringSubmatch(path)[1])
	if err != nil {
		log.Fatal("File couldn't be found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if regex.MatchString(line) {
			go readFromFile(line, regex, pattern, c, wg)
		} else {
			c <- strings.Count(line, *pattern)
		}
	}
}

func main() {
	initialFile := flag.String("initialFile", "./plikA.txt", "Path to the first file.")
	pattern := flag.String("pattern", "", "Pattern that should be counted by the program.")

	flag.Parse()

	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		readFromFile(fmt.Sprintf(`\input{%s}`, *initialFile), regex, pattern, c, &wg)

		wg.Wait()
		close(c)
	}()

	result := 0
	for value := range c {
		result += value
	}

	fmt.Println(result)
}
