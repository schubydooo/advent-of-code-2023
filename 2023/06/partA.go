package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type raceResult struct {
	time     int
	distance int
}

func main() {
	fmt.Println("Welcome to Advent of Code Day 6")

	// Variables
	solution := 0
	var results []raceResult

	// Read input file
	// schematic, err := readInputFile("sample.txt")
	schematic, err := readInputFile("input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	// Iterate through each row
	for y, row := range schematic {
		fmt.Println("row ", y, " - ", row)
		parsedString := strings.Split(row, " ")
		lineDetails := parsedString[0]

		// Iterate through race results for row
		i := 0
		for _, result := range parsedString {
			if result != "" && result != "Distance:" && result != "Time:" {
				fmt.Println("parsed string: ", result)
				if lineDetails == "Time:" {
					results = append(results, raceResult{
						time:     mustAtoI(result),
						distance: 0,
					})
					i++
				} else {
					updatedResult := results[i]
					updatedResult.distance = mustAtoI(result)
					results[i] = updatedResult
					i++
				}
			}
		}
	}

	// Calculate the longest race
	waysToWinRace := make([]int, len(results))
	for i, result := range results {
		waysToWinRace[i] = 0

		// For our boat, calculate distance per hold duration
		for buttonHoldTime := 0; buttonHoldTime < result.time; buttonHoldTime++ {
			speed := buttonHoldTime
			timeToTravel := result.time - buttonHoldTime
			distanceTraveled := speed * timeToTravel

			if distanceTraveled > result.distance {
				waysToWinRace[i]++
			}

		}

	}
	fmt.Println("Results: ", results)
	fmt.Println("WaysToWinRace: ", waysToWinRace)

	// Calculate total permutations of winning
	solution = 1
	for _, ways := range waysToWinRace {
		solution = solution * ways
	}
	fmt.Println("Solution: ", solution)
}

// Helper function
// Convert string to int
func mustAtoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Could not convert string to int: ", s)
		panic("yikes")
	}

	return i
}

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
