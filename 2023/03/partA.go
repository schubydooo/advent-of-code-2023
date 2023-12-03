package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type coordinate struct {
	X int
	Y int
}

func main() {
	fmt.Println("Welcome to Advent of Code Day 2")

	// Variables
	numbers := make(map[coordinate]int)
	solution := 0

	// Read input file
	schematic, err := readInputFile("input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	// Helper methods
	isSymbol := func(r rune) bool {
		return !unicode.IsDigit(r) && r != '.'
	}

	expandCurrentNumber := func(loc coordinate) (coordinate, int) {
		numStr := ""
		startLoc := coordinate{
			X: loc.X,
			Y: loc.Y,
		}

		// Find start of number
		for unicode.IsDigit(rune(schematic[startLoc.Y][startLoc.X-1])) {
			startLoc.X -= 1

			// Break at start of line
			if startLoc.X == 0 {
				break
			}
		}

		// Find end of number
		i := startLoc.X
		for unicode.IsDigit(rune(schematic[startLoc.Y][i])) {
			numStr += string(schematic[startLoc.Y][i])
			i++

			// Break if at end of line
			if i >= len(schematic[startLoc.Y]) {
				break
			}
		}

		// Convert to int
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("Couldn't convert number")
		}

		return startLoc, num
	}

	addAdjacentNumbersToMap := func(loc coordinate) {
		directions := []struct{ x, y int }{
			{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1},
		}

		for _, dir := range directions {
			inspect_x, inspect_y := loc.X+dir.x, loc.Y+dir.y

			// Only check in-bounds values, assume square grid
			if inspect_x >= 0 && inspect_y >= 0 && inspect_x < len(schematic[loc.Y]) && inspect_y < len(schematic) {
				inspect_val := rune(schematic[inspect_y][inspect_x])
				fmt.Println("        -> looking at ch's: ", string(inspect_val))

				if unicode.IsDigit(inspect_val) {
					fmt.Println("            FOUND A NUMBER")
					startLoc, num := expandCurrentNumber(coordinate{X: inspect_x, Y: inspect_y})
					numbers[startLoc] = num // Add to map
				}
			}
		}
	}

	// Iterate through each row
	for y, row := range schematic {
		fmt.Println("row ", y, " - ", row)

		// Find symbols and look for numbers in all directions
		for x, ch := range row {
			fmt.Println("    ch ", x, " - ", string(ch))
			if isSymbol(ch) {
				// If a number is found, find the start and end of the number
				fmt.Println("        Found a symbol")
				addAdjacentNumbersToMap(coordinate{X: x, Y: y})

			}

		}
	}

	// Sum engine parts
	for _, num := range numbers {
		solution += num
	}
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Solution: ", solution)
}

// Helper function
func readInputFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
