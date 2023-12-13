package main

import (
	"AOC23/utils"
	"math/rand"
	"strconv"
	"strings"
)

func main1() {
	fileName := "./day7/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	utils.Time(func() {
		part1(lineReader)
	})
	/*utils.Time(func() {
		part2(lineReader)
	})*/
}

func part1(lines []string) {

	handToBid, hands := parseToBidMap(lines)

	println("Finished parsing")

	hands = quickSort(hands, isHandLess)

	println("Finished sorting")

	var total int
	for i, hand := range hands {
		total += handToBid[hand.Values] * (i + 1)
	}

	println("Total: ", total)

}

func parseToBidMap(lines []string) (map[string]int, []Hand) {
	handToBidMap := make(map[string]int)
	hands := make([]Hand, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, ParseHand(hand))
		handToBidMap[hand] = bid
	}

	return handToBidMap, hands
}

func quickSort(hands []Hand, comparer func(Hand, Hand) bool) []Hand {
	if len(hands) < 2 {
		return hands
	}

	left, right := 0, len(hands)-1

	pivot := rand.Int() % len(hands)

	hands[pivot], hands[right] = hands[right], hands[pivot]

	for i := range hands {
		if comparer(hands[i], hands[right]) {
			hands[i], hands[left] = hands[left], hands[i]
			left++
		}
	}

	hands[left], hands[right] = hands[right], hands[left]

	quickSort(hands[:left], comparer)
	quickSort(hands[left+1:], comparer)

	return hands
}

type Hand struct {
	Type   int
	Values string
}

func isHandLess(hand, compare Hand) bool {
	if hand.Type < compare.Type {
		return true
	}

	if hand.Type > compare.Type {
		return false
	}

	if hand.Type == compare.Type {
		return isLessCompareByCard(hand.Values, compare.Values)
	}

	return false
}

func isLessCompareByCard(hand, compare string) bool {
	for i := 0; i < len(hand); i++ {
		if cardValue[string(hand[i])] == cardValue[string(compare[i])] {
			continue
		}

		if cardValue[string(hand[i])] < cardValue[string(compare[i])] {
			return true
		} else {
			return false
		}
	}

	return false
}

func ParseHand(hand string) Hand {

	handMap := make(map[string]int)
	for _, char := range hand {
		handMap[string(char)] += 1
	}

	var pairs int
	var triples int
	for _, value := range handMap {
		switch value {
		case 2:
			pairs += 1
		case 3:
			triples += 1
		case 4:
			return Hand{
				Type:   5,
				Values: hand,
			}
		case 5:
			return Hand{
				Type:   6,
				Values: hand,
			}
		}
	}

	if pairs == 1 && triples == 1 {
		return Hand{
			Type:   4,
			Values: hand,
		}
	}

	if pairs == 1 {
		return Hand{
			Type:   1,
			Values: hand,
		}
	}

	if pairs == 2 {
		return Hand{
			Type:   2,
			Values: hand,
		}
	}

	if triples == 1 {
		return Hand{
			Type:   3,
			Values: hand,
		}
	}

	return Hand{
		Type:   0,
		Values: hand,
	}
}

var cardValue = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}
