package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	str := getInputString()
	words := getWordsFromString(str)
	wordCount := len(words)

	fmt.Println(wordCount)
}

func getWordsFromString(str string) []string {
	return strings.Fields(str)
}

func getInputString() string {
	input := ""
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	return input
}
