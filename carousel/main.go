/*
Example of line filter
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
	os.Exit(1)
}

var (
	wg sync.WaitGroup
)

func worker(line string) {
	num, err := strconv.ParseInt(line, 10, 0)
	if err != nil {
		printError(err)
	}

	if num%2 == 0 {
		fmt.Fprintf(os.Stdout, "0: %d\n", num)
	} else {
		fmt.Fprintf(os.Stdout, "1: %d\n", num)
	}

	time.Sleep(time.Second)
	wg.Done()
}

func dispetcher(numbers chan string) {
	for {
		select {
		case number := <- numbers:
			log("Working on: " + number)
			go worker(number)
			wg.Add(1)
		}
	}
}

func log(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	max := 1000
	i := 0
	numbers := make(chan string)

	go dispetcher(numbers)

	for scanner.Scan() {
		select {
			case numbers <- scanner.Text():
				log("Send numbers: " + scanner.Text())
				if i > max {
					wg.Wait()
					os.Exit(0)
				} else {
					i++
				}
		}
	}
	if err := scanner.Err(); err != nil {
		printError(err)
	}
}
