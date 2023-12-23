package main

import (
	"bufio"
	"fmt"
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
	winningMatchCnt := make([]int, len(schematic))
	for y, row := range schematic {
		fmt.Println("row ", y, " - ", row)

		parsed_row := strings.Split(strings.Split(row, ":")[1], "|")

		// Extract numbers into slices
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
		winningMatchCnt[y] = numberMatchCnt

		fmt.Println("     Winning:", winning_numbers)
		fmt.Println("    Scratch:", scratch_numbers)
	}

	// Sum up how many scratch cards we have
	scratchCardCnt := make([]int, len(schematic))
	for cardNum := 0; cardNum < len(winningMatchCnt); cardNum++ {
		// Include the original card for the slot
		scratchCardCnt[cardNum] += 1

		// fmt.Println("Matched numbers for Card ", cardNum, ": ", winningMatchCnt[cardNum])

		// fmt.Println("    scratch card count: ", scratchCardCnt)
		// fmt.Println("    Looping through card ", cardNum, " ", scratchCardCnt[cardNum], " times")
		for i := 0; i < scratchCardCnt[cardNum]; i++ {
			// fmt.Println("    loop num ", i)
			for j := 0; j < winningMatchCnt[cardNum]; j++ {
				if j+cardNum+1 < len(winningMatchCnt) {
					// fmt.Println("        Add scratch ticket count to card", j+cardNum+1)
					scratchCardCnt[j+cardNum+1] += 1
					// time.Sleep(time.Second)
				}
			}
		}
		// fmt.Println("    scratch card count: ", scratchCardCnt)
		solution += scratchCardCnt[cardNum]
	}

	fmt.Println("winningMatchCnt: ", winningMatchCnt)
	fmt.Println("scratchCardCnt: ", scratchCardCnt)
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
