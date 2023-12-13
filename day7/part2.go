package main

import (
	"AOC23/utils"
	"slices"
	"strconv"
	"strings"
)

var cardValue2 = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
}

func main() {
	fileName := "./day7/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	utils.Time(func() {
		part2(lineReader)
	})
}

func part2(lines []string) {

	handToBid, hands := parseToBidMap2(lines)

	println("Finished parsing")

	hands = quickSort(hands, isHandLess2)

	println("Finished sorting")

	var total int
	for i, hand := range hands {
		total += handToBid[hand.Values] * (i + 1)
	}

	println("Total: ", total)

}

func isHandLess2(hand, compare Hand) bool {
	if hand.Type < compare.Type {
		return true
	}

	if hand.Type > compare.Type {
		return false
	}

	if hand.Type == compare.Type {
		return isLessCompareByCard2(hand.Values, compare.Values)
	}

	return false
}

func isLessCompareByCard2(hand, compare string) bool {
	for i := 0; i < len(hand); i++ {
		if cardValue2[string(hand[i])] == cardValue2[string(compare[i])] {
			continue
		}

		if cardValue2[string(hand[i])] < cardValue2[string(compare[i])] {
			return true
		} else {
			return false
		}
	}

	return false
}

func parseToBidMap2(lines []string) (map[string]int, []Hand) {
	handToBidMap := make(map[string]int)
	hands := make([]Hand, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, ParseHand2(hand))
		handToBidMap[hand] = bid
	}

	return handToBidMap, hands
}

func ParseHand2(hand string) Hand {

	var jokers int
	handMap := make(map[string]int)
	for _, char := range hand {
		if string(char) == "J" {
			jokers += 1
			continue
		}
		handMap[string(char)] += 1
	}

	frequencies := make([]int, 0)
	for _, value := range handMap {
		frequencies = append(frequencies, value)
	}

	slices.Sort(frequencies)

	var tp int
	for _, freq := range frequencies {
		switch freq {
		case 4:
			tp = 5
		case 5:
			tp = 6
		case 2:
			tp++
		case 3:
			tp += 3
		}
	}

	if jokers > 0 {
		tp = upgradeWithJokers(tp, jokers)
	}

	return Hand{
		Type:   tp,
		Values: hand,
	}

}

func upgradeWithJokers(handType int, jokers int) int {
	if handType == 5 {
		return 6
	}

	if handType == 4 || handType == 2 {
		return 4
	}

	if handType == 3 {
		if jokers == 1 {
			return 5
		} else {
			return 6
		}
	}

	if handType == 1 {
		if jokers == 1 {
			return 3
		} else if jokers == 2 {
			return 5
		} else if jokers == 3 {
			return 6
		}
	}

	if handType == 0 {
		if jokers == 1 {
			return 1
		} else if jokers == 2 {
			return 3
		} else if jokers == 3 {
			return 5
		} else if jokers == 4 || jokers == 5 {
			return 6
		}

	}

	panic("Invalid hand type")
}
