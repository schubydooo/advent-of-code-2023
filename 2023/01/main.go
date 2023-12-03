package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	fmt.Println("Welcome to Advent of Code Day 1")
	solvePart1()
	solvePart2()
}

// Part 2 of the puzzle
func solvePart2() {
	fmt.Println("Now we need to look at the letters to determine if the number is spelled out")
	// Import file
	calibration_values := readFile()

	// Parse each line for 1st and last int
	final_value := 0
	for _, value := range calibration_values {
		if value == "" {
			continue
		}
		og := value
		value = convertSpelledNumbers(value)
		first := getFirstNumber(value)
		second := getSecondNumber(value)

		calibration_val := (first * 10) + second
		// fmt.Println(value, " --> ", calibration_val)
		fmt.Println(og, " --> ", value, " --> ", calibration_val)
		final_value += calibration_val
	}
	fmt.Println("Part 2 - Final calibration value: ", final_value)
}

func convertSpelledNumbers(row string) string {
	// Preprocess string and convert spelled numbers to number chars
	// hacky approach to replace AND include edge strings to handle: eightwo -> 8wo
	row = strings.Replace(row, "one", "o1e", -1)
	row = strings.Replace(row, "two", "t2o", -1)
	row = strings.Replace(row, "three", "t3e", -1)
	row = strings.Replace(row, "four", "f4r", -1)
	row = strings.Replace(row, "five", "f5e", -1)
	row = strings.Replace(row, "six", "s6x", -1)
	row = strings.Replace(row, "seven", "s7n", -1)
	row = strings.Replace(row, "eight", "e8t", -1)
	row = strings.Replace(row, "nine", "n9e", -1)
	return row
}

// Part 1 of the puzzle
func solvePart1() {
	// Import file
	calibration_values := readFile()

	// Parse each line for 1st and last int
	final_value := 0
	for _, value := range calibration_values {
		if value == "" {
			continue
		}
		first := getFirstNumber(value)
		second := getSecondNumber(value)

		calibration_val := (first * 10) + second
		// fmt.Println(value, " --> ", calibration_val)
		final_value += calibration_val
	}

	fmt.Println("Part 1 - Final calibration value: ", final_value)
}

func parseRow(row string) int {
	for i := 0; i < len(row); i++ {
		char := string(row[i])
		if slices.Contains(digits, char) {
			i, err := strconv.Atoi(char)
			if err != nil {
				panic("Couldn't convert string to int")
			}

			return i
		}
	}
	panic("Didn't find an integer on this line")
}

func getFirstNumber(row string) int {
	return parseRow(row)
}

func getSecondNumber(row string) int {
	return parseRow(reverse(row))
}

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func readFile() []string {
	// Open the CSV file
	file, err := os.Open("calibration_values.csv")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	// Copy the CSV to a local variable
	data := make([]string, 1000) //1000 appears to be the expected input file len
	i := 0
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("Error reading row")
		}

		// Data validation
		if len(record) > 1 {
			panic("CSV has more than 1 value per row")
		}
		if len(record[0]) < 1 {
			panic("Line has no values")
		}

		data[i] = record[0]
		i++
	}

	return data
}
