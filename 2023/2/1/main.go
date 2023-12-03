package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var limits = map[string]int{
	"blue":  14,
	"red":   12,
	"green": 13,
}

type Set struct {
	id    int
	pulls []Pull
	bust  bool
}

type Pull struct {
	cubes  []Cube
	colors map[string]int
}

type Cube struct {
	color  string
	number int
}

func main() {
	sum := 0

	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		set := Set{}
		set.bust = false

		s := strings.Split(text, ": ")
		idStr, pullsStr := s[0], s[1]

		// get the set ID to sum later
		s = strings.Split(idStr, " ")
		set.id, _ = strconv.Atoi(s[1])

		// the elf pulls multiple cubes at a time
		p := strings.Split(pullsStr, "; ")
		for _, pullStr := range p {
			pull := Pull{}
			pull.colors = map[string]int{"blue": 0, "red": 0, "green": 0}

			cubes := strings.Split(pullStr, ", ")

			// Record all the cubes
			for _, cubeStr := range cubes {
				cube := Cube{}
				s = strings.Split(cubeStr, " ")
				cube.number, _ = strconv.Atoi(s[0])
				cube.color = s[1]
				pull.cubes = append(pull.cubes, cube)
			}

			// Add all cubes
			for _, cube := range pull.cubes {
				pull.colors[cube.color] += cube.number
			}

			// Compare totals to limits
			for color, total := range pull.colors {
				if total > limits[color] {
					set.bust = true // if any color is over the limit, flag the whole set
				}
			}
		}

		if !set.bust {
			sum += set.id // add all sets that dont go over limit
		}
	}

	fmt.Println(sum)
}
