package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText := scanner.Text()

		// find first digit
		firstDigit, forwardFound := traverseForward(inputText)
		if !forwardFound {
			fmt.Errorf("No digit found traversing forward in text: %s", inputText)
		}

		// find last digit
		lastDigit, backwardsFound := traverseBackwards(inputText)
		if !backwardsFound {
			fmt.Errorf("No digit found traversing forward in text: %s", inputText)
		}

		firstInt, err := strconv.Atoi(string(firstDigit))
		if err != nil {
			log.Fatal(err)
		}

		lastInt, err := strconv.Atoi(string(lastDigit))
		if err != nil {
			log.Fatal(err)
		}

		sum += (firstInt * 10) + lastInt

	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func traverseForward(text string) (rune, bool) {
	for _, n := range text {
		if unicode.IsDigit(n) {
			return n, true
		}
	}
	return 0, false
}

func traverseBackwards(text string) (rune, bool) {
	lenText := len(text)
	for i, _ := range text {
		if unicode.IsDigit(rune(text[lenText-(i+1)])) {
			return rune(text[lenText-(i+1)]), true
		}
	}
	return 0, false
}
