/*
Linux grep implementation in Go
Faster than linux grep for normal string matching.
Can be slightly slower for regex matching.

Usage:
⇒  go install gogrep.go
⇒  cat shakespeare.txt | gogrep "and"

Performance:
⇒  cat shakespeare.txt | time gogrep "and" > /dev/null
gogrep "and" > /dev/null  0.06s user 0.05s system 148% cpu 0.072 total

⇒  cat shakespeare.txt | time grep "and" > /dev/null
grep "and" > /dev/null  0.11s user 0.00s system 99% cpu 0.116 total

Regex:
⇒  cat shakespeare.txt | time gogrep "\Wand\W" > /dev/null
gogrep "\Wand\W" > /dev/null  0.25s user 0.06s system 119% cpu 0.253 total

⇒  cat shakespeare.txt | time grep "\Wand\W" > /dev/null
grep "\WCand\W" >   0.21s user 0.00s system 99% cpu 0.216 total

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var filterStr = os.Args[1]
var r, regexErr = regexp.Compile(filterStr)
var scanner = bufio.NewScanner(os.Stdin)

func grep() {
	line := scanner.Text()
	if r.MatchString(line) {
		fmt.Println(line)
	}
}

func main() {
	if regexErr != nil {
		log.Fatal(regexErr)
	}

	for scanner.Scan() {
		grep()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
