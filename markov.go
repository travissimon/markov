package main

import (
	"bufio"
	"log"
	"strings"
)


func main() {
	wordMap := make(map[string]string)

	const input = "Something has to change. Undeniable dilemma. Boredom's not a burden anyone should have to bear."


	// Create a word scanner
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	var deduped string
	for scanner.Scan() {
		nextWord := scanner.Text()
		found := false
		// de-dupe word instances
		if deduped, found = wordMap[nextWord]; !found {
			wordMap[nextWord] = nextWord
			deduped = nextWord
		}

		log.Printf("%s\n", deduped)
	}
}
