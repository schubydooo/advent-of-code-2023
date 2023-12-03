package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
    fmt.Println("Welcome to Advent of Code Day 1")

	// Import file 
	// calibration_values := readFile()
	readFile()

	// Parse each line for 1st and last int 
    // for _, line := range calibration_values {

	// 	// Data validation
	// 	if len(line[0]) < 1 {
	// 		fmt.Println("Line has no values")
	// 	}
	// 	fmt.Println(len(line[0]))
	// }


	// Sum returned lines
	// TODO

}

func parseRow(row string) int{
	return 1
}


func findFirstNum(row string) int {
	return 1
}

func readFile() {
	// Open the CSV file
	file, err := os.Open("calibration_values.csv")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	// Copy the CSV to a local variable
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading row: ", err)
			return 
		}

		fmt.Println(record)
	}


}