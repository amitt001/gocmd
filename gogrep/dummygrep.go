/*
Fancy but slower version of gogrep
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

var filterStr = os.Args[1]
var resCh = make(chan string)
var wg sync.WaitGroup
var r, regexErr = regexp.Compile(filterStr)

func grep() {
	for line := range resCh {
		if r.MatchString(line) {
			fmt.Println(line)
		}
		wg.Done()
	}
}

func main() {
	if regexErr != nil {
		log.Fatal(regexErr)
	}
	defer close(resCh)
	defer wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)
	go grep()

	for scanner.Scan() {
		wg.Add(1)
		resCh <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
