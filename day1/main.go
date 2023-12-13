package main

import (
	"AOC23/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var numbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {
	fileName := "./day1/input.txt"
	lineReader := utils.ReadFile(fileName)
	//part1(lineReader)
	part2(lineReader)

}

func part1(lines *bufio.Scanner) {
	var total int
	for lines.Scan() {
		line := lines.Text()

		var firstNumber string
		var secondNumber string

		for _, char := range line {
			if unicode.IsNumber(char) {
				if firstNumber == "" {
					firstNumber = string(char)
					secondNumber = string(char)

					continue
				}

				secondNumber = string(char)
			}
		}

		number, err := strconv.Atoi(firstNumber + secondNumber)
		if err != nil {
			panic(err)
		}

		total += number
	}

	fmt.Println("The total is: ", total)
}

func part2(lines *bufio.Scanner) {
	var total int
	for lines.Scan() {
		line := lines.Text()

		var firstNumber string
		var secondNumber string

		for i := 0; i < len(line); i++ {
			char := line[i]

			if unicode.IsNumber(rune(char)) {
				if firstNumber == "" {
					firstNumber = string(char)
					secondNumber = string(char)

					continue
				}

				secondNumber = string(char)
			}

			if unicode.IsLetter(rune(char)) {
				num, jump := decodeWrittenNumber(line[i:])
				if num == "" {
					i += jump
					continue
				}

				if firstNumber == "" {
					firstNumber = num
					secondNumber = num

					i += jump

					continue
				}

				secondNumber = num

				i += jump
			}
		}

		number, err := strconv.Atoi(firstNumber + secondNumber)
		if err != nil {
			panic(err)
		}

		total += number

	}

	fmt.Println("The total is: ", total)
}

func decodeWrittenNumber(text string) (string, int) {
	var textAccumulator string
	for i := 0; i < len(text); i++ {
		char := text[i]

		if unicode.IsNumber(rune(char)) {
			return "", i - 1
		}

		if unicode.IsLetter(rune(char)) {
			textAccumulator += string(char)
			for key, number := range numbers {
				if strings.Contains(textAccumulator, key) {
					return strconv.Itoa(number), i - 1
				}
			}
		}
	}

	return "", 0
}
