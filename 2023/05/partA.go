package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rangeLookup struct {
	sourceType  string
	destType    string
	sourceStart int
	destStart   int
	rangeLength int
}

func main() {
	fmt.Println("Welcome to Advent of Code Day 5")

	// Variables
	solution := 0
	var inputSeeds []int
	var lookups []rangeLookup
	var locations []int

	// Read input file
	// schematic, err := readInputFile("sample.txt")
	schematic, err := readInputFile("input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	// Iterate through each row
	var currentSource, currentDest string
	for y, row := range schematic {
		if strings.Contains(row, "seeds:") {
			line := strings.Split(row, " ")[1:]
			for _, seed := range line {
				inputSeeds = append(inputSeeds, mustAtoI(seed))
			}
			fmt.Println("Reading in input seeds: ", inputSeeds)
		} else if strings.Contains(row, "map:") {
			lookupString := strings.Split(row, " ")[0]
			currentSource = strings.Split(lookupString, "-")[0]
			currentDest = strings.Split(lookupString, "-")[2]
			fmt.Println("New source detected: ", currentSource)
			fmt.Println("New dest detected: ", currentDest)
		} else if row == "" {
			// ignore blank lines
			continue
		} else {
			lookupValues := strings.Split(row, " ")
			newLookup := rangeLookup{
				sourceType:  currentSource,
				destType:    currentDest,
				sourceStart: mustAtoI(lookupValues[1]),
				destStart:   mustAtoI(lookupValues[0]),
				rangeLength: mustAtoI(lookupValues[2]),
			}
			lookups = append(lookups, newLookup)
		}
		fmt.Println("row ", y, " - ", row)
	}

	fmt.Println("Lookups: ", lookups, "\n\n ")

	// for each seed calculate it's location
	for _, seed := range inputSeeds {
		fmt.Println("Evaluating seed: ", seed)
		loc := getLocationForSeed(seed, lookups)
		fmt.Println("    Returned location: ", loc)
		locations = append(locations, loc)
	}

	fmt.Println("Locations: ", locations)

	// Get closest location
	solution = slices.Min(locations)
	fmt.Println("Solution: ", solution)
}

func getLocationForSeed(seed int, lookups []rangeLookup) int {
	// Get soil for seed
	soil := convertValueToNextType(seed, lookups, "seed", "soil")
	fertilizer := convertValueToNextType(soil, lookups, "soil", "fertilizer")
	water := convertValueToNextType(fertilizer, lookups, "fertilizer", "water")
	light := convertValueToNextType(water, lookups, "water", "light")
	temperature := convertValueToNextType(light, lookups, "light", "temperature")
	humidity := convertValueToNextType(temperature, lookups, "temperature", "humidity")
	location := convertValueToNextType(humidity, lookups, "humidity", "location")
	return location
}

func convertValueToNextType(lookupValue int, lookups []rangeLookup, sourceType string, destType string) int {
	// Get soil for lookupValue
	for _, l := range lookups {
		if l.sourceType == sourceType && l.destType == destType {
			if lookupValue >= l.sourceStart && lookupValue < l.sourceStart+l.rangeLength {
				// lookupValue matches almanac.  Find new lookup type
				delta := lookupValue - l.sourceStart
				return l.destStart + delta
			}
		}
	}

	// If no match then it's a 1:1 mapping
	return lookupValue
}

// Helper functions
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
