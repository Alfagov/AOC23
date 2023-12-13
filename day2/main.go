package main

import (
	"AOC23/utils"
	"bufio"
	"strconv"
	"strings"
)

var (
	part1Limits = map[string]int{"blue": 14, "red": 12, "green": 13}
)

func main() {
	fileName := "./day2/input.txt"
	lineReader := utils.ReadFile(fileName)
	//part1(lineReader)
	/*utils.Time(func() {
		part1(lineReader)
	})*/
	utils.Time(func() {
		part2(lineReader)
	})
}

func part2(lines *bufio.Scanner) {
	var total int

	for lines.Scan() {
		line := lines.Text()

		gameSplit := strings.Split(line, ":")

		matches := strings.Split(gameSplit[1], ";")

		minimumMap := map[string]int{"blue": 0, "red": 0, "green": 0}

		for _, match := range matches {
			numColorPairs := strings.Split(match, ",")

			for _, numColorPair := range numColorPairs {
				numColorSplit := strings.Split(numColorPair, " ")
				num, err := strconv.Atoi(numColorSplit[1])
				if err != nil {
					panic(err)
				}
				color := numColorSplit[2]

				if minimumMap[color] < num {
					minimumMap[color] = num
				}
			}

		}

		var toAdd int = 1
		for _, val := range minimumMap {
			toAdd *= val
		}

		total += toAdd
	}

	println("The total is: ", total)
}

func part1(lines *bufio.Scanner) {
	var total int

MatchLine:
	for lines.Scan() {
		line := lines.Text()

		gameSplit := strings.Split(line, ":")

		gameHeader := gameSplit[0]
		gameId := gameHeader[5:]

		matches := strings.Split(gameSplit[1], ";")

		for _, match := range matches {
			numColorPairs := strings.Split(match, ",")

			for _, numColorPair := range numColorPairs {
				numColorSplit := strings.Split(numColorPair, " ")
				num, err := strconv.Atoi(numColorSplit[1])
				if err != nil {
					panic(err)
				}
				color := numColorSplit[2]

				if part1Limits[color] < num {
					continue MatchLine
				}
			}

		}

		gameIdInt, err := strconv.Atoi(gameId)
		if err != nil {
			panic(err)
		}
		total += gameIdInt

	}

	println("The total is: ", total)
}
