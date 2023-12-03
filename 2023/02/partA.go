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
	bag_contents := map[string]int{"red": 12, "green": 13, "blue": 14}

	// read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// parse file
	for sc.Scan() {
		row := sc.Text()
		sl := strings.Split(row, ":")
		game, _ := strconv.Atoi(strings.Split(sl[0], " ")[1])
		sets := strings.Split(sl[1], ";")

		// check each outcome of each set for impossible scenarios
		possible_game := true
		for _, s := range sets {
			for _, cubes_revealed := range strings.Split(s, ",") {
				cube_cnt, _ := strconv.Atoi(strings.Split(cubes_revealed, " ")[1])
				cube_color := strings.Split(cubes_revealed, " ")[2]
				if cube_cnt > bag_contents[cube_color] {
					possible_game = false
				}
			}

		}

		if possible_game {
			solution += game
		}

	}

	fmt.Println("Solution: ", solution)
}
