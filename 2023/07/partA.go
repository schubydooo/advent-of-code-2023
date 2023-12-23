package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	handNumber int
	cards      []int
	handType   int
	bid        int
}

func main() {
	fmt.Println("Welcome to Advent of Code Day 2")

	// Variables
	solution := 0
	var hands []hand

	// Read input file
	// schematic, err := readInputFile("sample.txt")
	schematic, err := readInputFile("input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	// Iterate through each row
	for y, row := range schematic {
		// Parse row
		fmt.Println("row ", y, " - ", row)
		cardsRaw := strings.Split(row, " ")[0]
		bid := mustAtoI(strings.Split(row, " ")[1])

		// Encode card information
		var cards []int
		for _, card := range cardsRaw {
			cards = append(cards, mustAtoI(camelCardsRuneToString(card)))
		}

		// Determine hand type
		handType := getHandType(cards)

		// Add to slice
		currentHand := hand{
			handNumber: y,
			cards:      cards,
			handType:   handType,
			bid:        bid,
		}
		hands = append(hands, currentHand)
		fmt.Println("    CurrentHand: ", currentHand)

	}

	// rank our hands strength
	fmt.Println("Time to rank our hands strength: ", hands)
	handTypes := make([]int, 0)
	for _, hand := range hands {
		fmt.Println("hand number ", hand.handNumber, "has type", hand.handType)
		handTypes = append(handTypes, hand.handType)
		fmt.Println("    sorted: ", handTypes)
		sort.Slice(handTypes, func(i, j int) bool {
			if handTypes[i] > handTypes[j] {
				return false
			}
			return true
		})
		fmt.Println("    sorted: ", handTypes)
	}

	fmt.Println("sorted hands: ", hands)
	sort.Slice(hands, func(i, j int) bool {
		fmt.Println("    comparing ", hands[i], " to ", hands[j])
		if hands[i].handType == hands[j].handType {
			// need to compare cards in order
			for cardNum := 0; cardNum < 5; cardNum++ {
				if hands[i].cards[cardNum] != hands[j].cards[cardNum] {
					return hands[i].cards[cardNum] < hands[j].cards[cardNum]
				}
			}
		}
		return hands[i].handType < hands[j].handType
	})
	fmt.Println("sorted hands: ", hands)

	// Calculate the total winnings
	for i, hand := range hands {
		solution += hand.bid * (i + 1)
	}
	fmt.Println("Solution: ", solution)
}

func isStrongerCard(baseCard int, compareCard int) bool {
	return true
}

func getHandType(cards []int) int {
	// Determine if hand is one of
	// 		7 Five of a kind, 	where all five cards have the same label: AAAAA
	// 		6 Four of a kind, 	where four cards have the same label and one card has a different label: AA8AA
	// 		5 Full house, 		where three cards have the same label, and the remaining two cards share a different label: 23332
	// 		4 Three of a kind,	where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	// 		3 Two pair, 			where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	// 		2 One pair, 			where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	// 		1 High card, 			where all cards' labels are distinct: 23456
	distinctCardCounts := make([]int, 15)
	for i := 0; i < 5; i++ {
		distinctCardCounts[cards[i]]++
	}
	fmt.Println("    Card representation: ", distinctCardCounts)

	if isFiveOfAKind(distinctCardCounts) {
		return 7
	} else if isFourOfAKind(distinctCardCounts) {
		return 6
	} else if isFullHouse(distinctCardCounts) {
		return 5
	} else if isThreeOfAKind(distinctCardCounts) {
		return 4
	} else if isTwoPair(distinctCardCounts) {
		return 3
	} else if isOnePair(distinctCardCounts) {
		return 2
	} else {
		// by default the type is high card
		return 1
	}
}

func isOnePair(cardRepr []int) bool {
	// Assume we have already checked for better hands
	if getPairCount(cardRepr) == 1 {
		return true
	}
	return false
}
func isTwoPair(cardRepr []int) bool {
	// Assume we have already checked for better hands
	if getPairCount(cardRepr) == 2 {
		return true
	}
	return false
}
func getPairCount(cardRepr []int) int {
	pairCount := 0
	for _, cardCount := range cardRepr {
		if cardCount == 2 {
			pairCount++
		}
	}

	return pairCount
}
func isThreeOfAKind(cardRepr []int) bool {
	// Assume we have already checked for better hands
	if slices.Max(cardRepr) == 3 {
		return true
	} else {
		return false
	}
}
func isFullHouse(cardRepr []int) bool {
	// Assume we have already checked for better hands
	if slices.Max(cardRepr) == 3 {
		for _, cardCount := range cardRepr {
			if cardCount == 2 {
				return true
			}
		}
	}
	return false
}

func isFourOfAKind(cardRepr []int) bool {
	// Assume we have already checked for 5kind
	if slices.Max(cardRepr) == 4 {
		return true
	} else {
		return false
	}
}

func isFiveOfAKind(cardRepr []int) bool {
	if slices.Max(cardRepr) == 5 {
		return true
	} else {
		return false
	}
}

func camelCardsRuneToString(r rune) string {
	s := string(r)
	if s == "T" {
		s = "10"
	} else if s == "J" {
		s = "11"
	} else if s == "Q" {
		s = "12"
	} else if s == "K" {
		s = "13"
	} else if s == "A" {
		s = "14"
	}
	return s
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
