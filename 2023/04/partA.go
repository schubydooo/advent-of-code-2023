package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Welcome to Advent of Code Day 2")

	// Variables
	solution := 0

	// Read input file
	// schematic, err := readInputFile("sample.txt")
	schematic, err := readInputFile("input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	// Iterate through each row
	for y, row := range schematic {
		fmt.Println("row ", y, " - ", row)

		parsed_row := strings.Split(strings.Split(row, ":")[1], "|")
		// Extract winning numbers into slice
		winning_numbers := strings.Split(parsed_row[0], " ")
		scratch_numbers := strings.Split(parsed_row[1], " ")

		numberMatchCnt := 0
		for _, myNum := range scratch_numbers {

			// If the value is not blank
			if myNum != "" {
				isWinningNumber := slices.Contains(winning_numbers, myNum)
				fmt.Println("        ", myNum, " - ", isWinningNumber)
				if isWinningNumber {
					numberMatchCnt += 1
				}
			}
		}

		if numberMatchCnt > 2 {
			solution += int(math.Pow(2, float64(numberMatchCnt-1)))
		} else {
			solution += numberMatchCnt
		}

		fmt.Println("     Winning:", winning_numbers)
		fmt.Println("    Scratch:", scratch_numbers)
	}

	// fmt.Println("Numbers: ", numbers)
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
