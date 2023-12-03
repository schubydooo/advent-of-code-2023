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
	calibration_values := readFile()

	// Parse each line for 1st and last int 
	fmt.Println(calibration_values)

	// Sum returned lines
	// TODO

}

func parseRow(row string) int{
	return 1
}


func findFirstNum(row string) int {
	return 1
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
		i ++
	}

	return data


}