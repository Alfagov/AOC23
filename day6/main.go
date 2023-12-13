package main

import (
	"AOC23/utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := "./day6/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	utils.Time(func() {
		part1(lineReader)
	})
	utils.Time(func() {
		part2(lineReader)
	})
}

type Race struct {
	Time     int
	Distance int
}

func part2(lines []string) {
	timeNumbers := strings.Split(lines[0], ":")
	distanceNumbers := strings.Split(lines[1], ":")

	timeString := strings.Replace(timeNumbers[1], " ", "", -1)
	distanceString := strings.Replace(distanceNumbers[1], " ", "", -1)

	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)

	println("Time: ", time)
	println("Distance: ", distance)

	var speed = 0
	var wins = 0
	for acc := 1; acc < time; acc++ {
		speed += 1
		travelled := (time - acc) * speed
		if travelled > distance {
			wins += 1
		}
	}

	println("Wins: ", wins)
}

func part1(lines []string) {
	timeNumbers := strings.Split(lines[0], ":")
	distanceNumbers := strings.Split(lines[1], ":")

	timeNumbersList := strings.Split(timeNumbers[1], " ")
	distanceNumbersList := strings.Split(distanceNumbers[1], " ")

	timeNumbersList = slices.DeleteFunc(timeNumbersList, utils.RemoveEmpty)
	distanceNumbersList = slices.DeleteFunc(distanceNumbersList, utils.RemoveEmpty)

	races := make([]Race, 0)
	for i := 0; i < len(timeNumbersList); i++ {

		time, _ := strconv.Atoi(timeNumbersList[i])
		distance, _ := strconv.Atoi(distanceNumbersList[i])

		races = append(races, Race{
			Time:     time,
			Distance: distance,
		})
	}

	var total = 1

	for _, race := range races {
		var wins int
		var speed = 0
		for acc := 1; acc < race.Time; acc++ {
			speed += 1
			travelled := (race.Time - acc) * speed
			if travelled > race.Distance {
				wins += 1
			}
		}

		total *= wins
	}

	println("Total wins: ", total)

}
