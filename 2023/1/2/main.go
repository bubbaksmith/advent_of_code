package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var wordCombos = map[string][2]int{
	"oneight":   {1, 8},
	"twone":     {2, 1},
	"threeight": {3, 8},
	"fiveight":  {5, 8},
	"nineight":  {9, 8},
	"eightwo":   {8, 2},
	"sevenine":  {7, 9},
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		pattern := regexp.MustCompile(`\d|oneight|twone|threeight|fiveight|nineight|eightwo|sevenine|one|two|three|four|five|six|seven|eight|nine`)
		allmatches := pattern.FindAllString(text, -1)

		var matches = [2]string{
			allmatches[0],
			allmatches[len(allmatches)-1],
		}

		var firstAndLast = [2]int{}

		// check if number
		for i, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				// check if a word combo
				if _, ok := wordCombos[match]; ok {
					num = wordCombos[match][i] // let the index dictate which version of the combo to return
				} else { // its just a word
					num = wordToNumber[match]
				}
			}
			firstAndLast[i] = num
		}

		first, last := firstAndLast[0], firstAndLast[1]

		sum += (first * 10) + last
	}

	fmt.Println(sum)
}
