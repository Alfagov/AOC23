package main

import (
	"AOC23/utils"
	"strconv"
	"strings"
)

func main() {
	fileName := "./day9/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	utils.Time(func() {
		part1(lineReader)
	})
	utils.Time(func() {
		part2(lineReader)
	})
}

func part1(lineReader []string) {
	total := 0
	for _, line := range lineReader {
		values := strings.Split(line, " ")
		valueList := make([]int, 0)

		for _, value := range values {
			num, _ := strconv.Atoi(value)
			valueList = append(valueList, num)
		}

		total += manageLine(valueList)
	}

	println("Part 1: ", total)
}

func part2(lineReader []string) {
	total := 0
	for _, line := range lineReader {
		values := strings.Split(line, " ")
		valueList := make([]int, 0)

		for _, value := range values {
			num, _ := strconv.Atoi(value)
			valueList = append(valueList, num)
		}

		total += manageLine2(valueList)
	}

	println("Part 2: ", total)
}

func manageLine(line []int) int {

	newList := make([]int, 0)
	for i := 1; i < len(line); i++ {
		newList = append(newList, line[i]-line[i-1])
	}

	zeroCount := 0
	for _, value := range newList {
		if value != 0 {
			break
		}
		zeroCount++
	}

	if zeroCount == len(newList) {
		return line[len(line)-1]
	}

	under := manageLine(newList)
	return line[len(line)-1] + under

}

func manageLine2(line []int) int {

	newList := make([]int, 0)
	for i := 1; i < len(line); i++ {
		newList = append(newList, line[i]-line[i-1])
	}

	zeroCount := 0
	for _, value := range newList {
		if value != 0 {
			break
		}
		zeroCount++
	}

	if zeroCount == len(newList) {
		return line[0]
	}

	under := manageLine2(newList)
	return line[0] - under

}
