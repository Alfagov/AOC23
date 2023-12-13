package main

import (
	"AOC23/utils"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileName := "./day5/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	/*utils.Time(func() {
		part1(lineReader)
	})*/
	utils.Time(func() {
		part2(lineReader)
	})
}

type MinMax struct {
	Min, Max int
}

type Target struct {
	SourceValues, TargetValues []MinMax
	Targ                       string
}

func part2(lines []string) {

	maps, seeds := parse2(lines)

	println("Finished parsing")

	//Find min recursiveSearch result
	minimum := 999999999999999999
	for _, seed := range seeds {

		for i := 0; i < seed[1]; i++ {
			value := recursiveSearch(maps, "seed", seed[0]+i)
			//println("Seed: ", seed)
			//println("Steps: ", value)

			if value < minimum {
				minimum = value
			}
		}

	}

	println("Minimum steps part 2: ", minimum)

}

func part1(lines []string) {

	maps, seeds := parse(lines)

	println("Finished parsing")

	//Find min recursiveSearch result
	minimum := 999999999999999999
	for _, seed := range seeds {

		value := recursiveSearch(maps, "seed", seed)
		println("Seed: ", seed)
		println("Steps: ", value)

		if value < minimum {
			minimum = value
		}
	}

	println("Minimum steps: ", minimum)

}

func recursiveSearch(maps map[string]Target, source string, value int) int {

	if _, ok := maps[source]; !ok {
		return value
	}

	target := maps[source]

	for i, sourceMinMax := range target.SourceValues {
		if sourceMinMax.Min <= value && value <= sourceMinMax.Max {
			targetMinMax := target.TargetValues[i]
			value = targetMinMax.Min + (value - sourceMinMax.Min)
			break
		}
	}

	return recursiveSearch(maps, target.Targ, value)
}

func parse2(lines []string) (map[string]Target, [][]int) {
	output := make([][]int, 0)
	var source, target string
	mapping := make(map[string]Target)
	for i := 0; i < len(lines); i++ {

		line := lines[i]

		if i == 0 {
			splits := strings.Split(line, ":")

			var divided [][]string

			chunkSize := 2
			seedsList := strings.Split(splits[1], " ")

			for i := 1; i <= len(seedsList); i += chunkSize {
				end := i + chunkSize

				if end > len(seedsList) {
					break
				}

				divided = append(divided, seedsList[i:end])
			}

			for _, d := range divided {
				seedStart, _ := strconv.Atoi(d[0])
				seedLenght, _ := strconv.Atoi(d[1])

				output = append(output, []int{seedStart, seedLenght})
			}

			continue
		}

		if line == "" {
			source = ""
			target = ""
			continue
		}

		if unicode.IsLetter(rune(line[0])) {
			mapHeader := strings.Split(line, " ")[0]

			source = strings.Split(mapHeader, "-")[0]
			target = strings.Split(mapHeader, "-")[2]
			continue
		}

		if unicode.IsDigit(rune(line[0])) {
			numbers := strings.Split(line, " ")
			destination, _ := strconv.Atoi(numbers[0])
			sources, _ := strconv.Atoi(numbers[1])
			steps, _ := strconv.Atoi(numbers[2])

			//m := generateMap(destination, sources, steps)

			if _, ok := mapping[source]; !ok {
				mapping[source] = Target{
					SourceValues: make([]MinMax, 0),
					TargetValues: make([]MinMax, 0),
					Targ:         target,
				}
			}

			sourceMinMax := MinMax{
				Min: sources,
				Max: sources + steps,
			}

			targetMinMax := MinMax{
				Min: destination,
				Max: destination + steps,
			}

			mapping[source] = Target{
				SourceValues: append(mapping[source].SourceValues, sourceMinMax),
				TargetValues: append(mapping[source].TargetValues, targetMinMax),
				Targ:         target,
			}

		}

	}

	return mapping, output
}

func parse(lines []string) (map[string]Target, []int) {
	seeds := make([]int, 0)
	var source, target string
	mapping := make(map[string]Target)
	for i := 0; i < len(lines); i++ {

		line := lines[i]

		if i == 0 {
			splits := strings.Split(line, ":")
			for _, n := range strings.Split(splits[1], " ") {
				if n == "" {
					continue
				}

				number, _ := strconv.Atoi(n)
				seeds = append(seeds, number)
			}

			continue
		}

		if line == "" {
			source = ""
			target = ""
			continue
		}

		if unicode.IsLetter(rune(line[0])) {
			mapHeader := strings.Split(line, " ")[0]

			source = strings.Split(mapHeader, "-")[0]
			target = strings.Split(mapHeader, "-")[2]
			continue
		}

		if unicode.IsDigit(rune(line[0])) {
			numbers := strings.Split(line, " ")
			destination, _ := strconv.Atoi(numbers[0])
			sources, _ := strconv.Atoi(numbers[1])
			steps, _ := strconv.Atoi(numbers[2])

			//m := generateMap(destination, sources, steps)

			if _, ok := mapping[source]; !ok {
				mapping[source] = Target{
					SourceValues: make([]MinMax, 0),
					TargetValues: make([]MinMax, 0),
					Targ:         target,
				}
			}

			sourceMinMax := MinMax{
				Min: sources,
				Max: sources + steps,
			}

			targetMinMax := MinMax{
				Min: destination,
				Max: destination + steps,
			}

			mapping[source] = Target{
				SourceValues: append(mapping[source].SourceValues, sourceMinMax),
				TargetValues: append(mapping[source].TargetValues, targetMinMax),
				Targ:         target,
			}

		}

	}

	return mapping, seeds
}

func generateMap(target, source, steps int) map[int]int {
	m := make(map[int]int)

	for i := 0; i < steps; i++ {
		m[source+i] = target + i
	}

	return m
}
