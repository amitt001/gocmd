/*
GO implementation of cat
*/

package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"os"
)

var source = os.Args[1]

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func main() {
	_, err := exists(source)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(source) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
}
