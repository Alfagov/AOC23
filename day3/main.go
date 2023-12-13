package main

import (
	"AOC23/utils"
	"log"
	"strconv"
	"unicode"
)

func main() {
	fileName := "./day3/input.txt"
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

type Point struct {
	Line, Row int
}

type ConfigurationMachine struct {
	lines             []string
	numberAccumulator string
	symbolMap         map[string][]Point
	numberMap         map[Point]int
	positionToNumbers map[Point][]int
}

func part1(lines []string) {
	config := &ConfigurationMachine{
		lines:             lines,
		numberAccumulator: "",
		symbolMap:         make(map[string][]Point),
		numberMap:         make(map[Point]int),
		positionToNumbers: make(map[Point][]int),
	}

	config.Parse()
	config.SumAdj()

}

func part2(lines []string) {
	config := &ConfigurationMachine{
		lines:             lines,
		numberAccumulator: "",
		symbolMap:         make(map[string][]Point),
		numberMap:         make(map[Point]int),
		positionToNumbers: make(map[Point][]int),
	}

	config.Parse()
	config.SumAdjP2()

}

func (cfg *ConfigurationMachine) SumAdjP2() {
	var total int

	for _, point := range cfg.symbolMap["*"] {
		numbers := cfg.positionToNumbers[point]

		if len(numbers) != 2 {
			continue
		}

		ratio := numbers[0] * numbers[1]

		total += ratio
	}

	log.Println("The total is: ", total)
}

func (cfg *ConfigurationMachine) SumAdj() {
	var total int
	for _, points := range cfg.symbolMap {
		for _, point := range points {
			numbers := cfg.positionToNumbers[point]
			var tmp int
			for _, number := range numbers {
				tmp += number
			}

			total += tmp
		}
	}

	log.Println("The total is: ", total)
}

func (cfg *ConfigurationMachine) Parse() {

	for lineIdx, line := range cfg.lines {
		for rowIdx, char := range line {

			if unicode.IsDigit(char) {
				cfg.numberAccumulator += string(char)

				if rowIdx == len(line)-1 {
					cfg.ParseAccumulator(lineIdx, rowIdx, false)
				}
				continue
			}

			if char == '.' {
				cfg.ParseAccumulator(lineIdx, rowIdx, true)
				continue
			}

			if cfg.symbolMap[string(char)] == nil {
				cfg.symbolMap[string(char)] = make([]Point, 0)
			}

			cfg.symbolMap[string(char)] = append(cfg.symbolMap[string(char)], Point{Line: lineIdx, Row: rowIdx})

			cfg.ParseAccumulator(lineIdx, rowIdx, false)

		}
	}
}

func (cfg *ConfigurationMachine) ParseAccumulator(lineIdx, rowIdx int, point bool) {
	if cfg.numberAccumulator != "" {
		number, err := strconv.Atoi(cfg.numberAccumulator)
		if err != nil {
			panic(err)
		}

		for i := 0; i <= len(cfg.numberAccumulator)+1; i++ {
			upPoint := Point{Line: lineIdx - 1, Row: rowIdx - i}
			downPoint := Point{Line: lineIdx + 1, Row: rowIdx - i}

			if cfg.positionToNumbers[upPoint] == nil {
				cfg.positionToNumbers[upPoint] = make([]int, 0)
			}

			if cfg.positionToNumbers[downPoint] == nil {
				cfg.positionToNumbers[downPoint] = make([]int, 0)
			}

			cfg.positionToNumbers[upPoint] = append(cfg.positionToNumbers[upPoint], number)
			cfg.positionToNumbers[downPoint] = append(cfg.positionToNumbers[downPoint], number)
		}

		beforePoint := Point{Line: lineIdx, Row: rowIdx - len(cfg.numberAccumulator) - 1}
		thisPoint := Point{Line: lineIdx, Row: rowIdx}

		if cfg.positionToNumbers[thisPoint] == nil {
			cfg.positionToNumbers[thisPoint] = make([]int, 0)
		}

		cfg.positionToNumbers[thisPoint] =
			append(cfg.positionToNumbers[thisPoint], number)

		if cfg.positionToNumbers[beforePoint] == nil {
			cfg.positionToNumbers[beforePoint] = make([]int, 0)
		}

		cfg.positionToNumbers[beforePoint] =
			append(cfg.positionToNumbers[beforePoint], number)

		cfg.numberAccumulator = ""
	}
}
