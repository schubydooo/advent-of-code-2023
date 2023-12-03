package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Advent of Code Day 2")

	// variables
	solution := 0

	// read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// parse file
	for sc.Scan() {
		row := sc.Text()
		sl := strings.Split(row, ":")
		sets := strings.Split(sl[1], ";")

		// check each outcome for minimum required cubes
		min_cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, s := range sets {
			for _, cubes_revealed := range strings.Split(s, ",") {
				cube_cnt, _ := strconv.Atoi(strings.Split(cubes_revealed, " ")[1])
				cube_color := strings.Split(cubes_revealed, " ")[2]

				min_cubes[cube_color] = max(cube_cnt, min_cubes[cube_color])
			}

		}

		solution += min_cubes["red"] * min_cubes["green"] * min_cubes["blue"]
	}

	fmt.Println("Solution: ", solution)

}
