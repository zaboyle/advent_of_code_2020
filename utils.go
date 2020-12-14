package main

import(
	"os"
	"io"
	"bufio"
	"log"
)

func contains(slice []string, str string) bool {
	for _, elt := range slice {
		if elt == str {
			return true
		}
	}
	return false
}

func getFileReader(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file);
	return r
}