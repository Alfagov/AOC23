package main

import (
	"AOC23/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := "./day4/input.txt"
	lineReader := utils.ReadFileLines(fileName)
	//part1(lineReader)
	/*utils.Time(func() {
		part1(lineReader)
	})*/
	utils.Time(func() {
		part1(lineReader)
	})
	utils.Time(func() {
		part2(lineReader)
	})
}

func part1(lines []string) {
	var total int
	for _, line := range lines {

		winningCount := ParseLineToCards(line)

		points := math.Pow(2, float64(winningCount-1))

		total += int(points)
	}

	println("Total points is: ", total)
}

type ScratchCards struct {
	WinningCount int
}

func part2(lines []string) {

	scratchCardsToWinningCount := make([]int, 0)
	for _, line := range lines {
		winningCount := ParseLineToCards(line)
		scratchCardsToWinningCount = append(scratchCardsToWinningCount, winningCount)
	}

	println("Total points is: ", recursiveCount(scratchCardsToWinningCount, len(scratchCardsToWinningCount)))
}

func recursiveCount(scratchCards []int, num int) int {
	total := 0
	for i := 0; i < num; i++ {
		c := scratchCards[i]
		if c == 0 {
			total++
			continue
		}

		total++
		total += recursiveCount(scratchCards[i+1:], c)
	}

	return total
}

func ParseLineToCards(line string) int {
	headerBodySplit := strings.Split(line, ":")

	cardsSplit := strings.Split(headerBodySplit[1], "|")

	winningCardsString := strings.Split(cardsSplit[0], " ")
	winningCards := make([]int, 0)
	for _, card := range winningCardsString {
		if card == "" {
			continue
		}
		num, _ := strconv.Atoi(card)
		winningCards = append(winningCards, num)
	}

	playerCardsString := strings.Split(cardsSplit[1], " ")
	playerCards := make([]int, 0)
	for _, card := range playerCardsString {
		if card == "" {
			continue
		}
		num, _ := strconv.Atoi(card)
		playerCards = append(playerCards, num)
	}

	winningCount := 0
	for _, card := range playerCards {
		if slices.Index(winningCards, card) != -1 {
			winningCount++
		}
	}

	return winningCount
}
