package utils

import (
	"bufio"
	"os"
	"time"
)

func Time(action func()) {
	start := time.Now()
	action()
	elapsed := time.Now().Sub(start).Microseconds()

	println("Time taken: ", elapsed, "µs")
}

func ReadFile(file string) *bufio.Scanner {
	readFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func ReadFileLines(file string) []string {
	fileScanner := ReadFile(file)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func RemoveEmpty(n string) bool {
	if n == "" {
		return true
	} else {
		return false
	}
}
