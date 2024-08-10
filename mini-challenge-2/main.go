package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}

	sentence = strings.TrimSpace(sentence)
	calculateWord(sentence)
}

func calculateWord(word string) map[string]int {
	var counter = make(map[string]int)
	for _, letter := range word {
		fmt.Println(string(letter))

		if _, exists := counter[string(letter)]; exists {
			counter[string(letter)]++
		} else {
			counter[string(letter)] = 1
		}
	}

	fmt.Println(counter)

	return counter
}
