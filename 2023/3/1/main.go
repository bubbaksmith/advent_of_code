package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x     int
	y     int
	empty bool
}

func main() {
	var fullDiagram = [141][]string{}
	var sum = 0

	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	lineIndex := 0
	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Printf("Read in text: %s\n", text)
		lineArr := strings.Split(text, "")
		fullDiagram[lineIndex] = append(fullDiagram[lineIndex], lineArr...)
		lineIndex += 1
	}

	var start Coordinate
	start.empty = true
	var end Coordinate
	end.empty = true

	for i, line := range fullDiagram {
		for j, char := range line {
			// find the first number
			if _, err := strconv.Atoi(char); start.empty && err == nil {
				start.x = j
				start.y = i
				start.empty = false
			}

			// find the end of number sequence
			_, err := strconv.Atoi(char)
			if !start.empty && err != nil || j == len(line)-1 {
				end.x = j
				end.y = i

				start.empty = true

				numberArr := line[start.x:end.x]

				// if is the last element of a line, we need to append it with our index logic
				if j == len(line)-1 {
					numberArr = line[start.x:]
				}

				numberStr := strings.Join(numberArr, "")
				num, _ := strconv.Atoi(numberStr)

				fmt.Printf("num: '%d'\n", num)

				if isPartNumber(start.x, end.x, start.y, fullDiagram) {
					sum += num
				}

			}

		}
	}

	fmt.Println(sum)
}

func isPartNumber(xStart int, xEnd int, currentLineIdx int, fullDiagram [141][]string) bool {
	var previousLineIdx int
	if currentLineIdx == 0 {
		previousLineIdx = 0
	} else {
		previousLineIdx = currentLineIdx - 1
	}

	var nextLineIdx int
	if fullDiagram[currentLineIdx+1] == nil {
		nextLineIdx = currentLineIdx
	} else {
		nextLineIdx = currentLineIdx + 1
	}

	var sideStart int
	if xStart == 0 {
		sideStart = 0
	} else {
		sideStart = xStart - 1
	}

	var sideEnd int
	if xEnd == len(fullDiagram[currentLineIdx])-1 {
		sideEnd = xEnd
	} else {
		sideEnd = xEnd + 1
	}

	upperText := slices.Clone(fullDiagram[previousLineIdx][sideStart:sideEnd])
	lowerText := slices.Clone(fullDiagram[nextLineIdx][sideStart:sideEnd])
	leftSide := fullDiagram[currentLineIdx][sideStart]
	rightSide := fullDiagram[currentLineIdx][sideEnd]

	if xEnd == len(fullDiagram[currentLineIdx])-1 {
		upperText = fullDiagram[previousLineIdx][sideStart:]
		lowerText = fullDiagram[nextLineIdx][sideStart:]
	}

	combinedNeighbors := append(upperText, lowerText...)
	combinedNeighbors = append(combinedNeighbors, leftSide, rightSide)

	pattern := `[0-9\.]`
	regexpPattern := regexp.MustCompile(pattern)

	fmt.Println(upperText)
	fmt.Printf("[%s       %s]\n", leftSide, rightSide)
	fmt.Println(lowerText)

	fmt.Println(combinedNeighbors)

	// check if any element doesn't match the regular expression
	for _, element := range combinedNeighbors {
		if !regexpPattern.MatchString(element) {
			fmt.Printf("Found symbol '%s'\n", element)
			return true
		}
	}

	return false
}
