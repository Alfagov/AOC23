package main

import (
	"AOC23/utils"
	"slices"
	"strings"
)

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Val   []string
	Edges map[string]*Edge
}

type Edge struct {
	Vertex *Vertex
}

func (g *Graph) AddVertex(key string, val []string) {
	if _, ok := g.Vertices[key]; ok {
		if g.Vertices[key].Val == nil && val != nil {
			g.Vertices[key].Val = val
		}
		return
	}
	g.Vertices[key] = &Vertex{Val: val, Edges: map[string]*Edge{}}
}

func (g *Graph) AddEdge(from, to string) {
	if _, ok := g.Vertices[from]; !ok {
		return
	}

	if _, ok := g.Vertices[to]; !ok {
		return
	}

	g.Vertices[from].Edges[to] = &Edge{Vertex: g.Vertices[to]}
}

func main() {
	fileName := "./day8/input.txt"
	lineReader := utils.ReadFileLines(fileName)

	/*utils.Time(func() {
		part1(lineReader)
	})*/
	utils.Time(func() {
		part2(lineReader)
	})
}

type Node struct {
	Left, Right string
}

func part2(lineReader []string) {
	instructions := lineReader[0]
	usefulLines := lineReader[2:]

	nodes := make(map[string]Node)

	var starts []string
	var ends []string

	for _, nodeLine := range usefulLines {
		fields := strings.Fields(nodeLine)

		if fields[0][2] == 'A' {
			starts = append(starts, fields[0])
		}

		if fields[0][2] == 'Z' {
			ends = append(ends, fields[0])
		}

		left, right := fields[2], fields[3]
		nodes[fields[0]] = Node{Left: left[1:4], Right: right[:len(right)-1]}
	}

	var paths []int

	for _, start := range starts {
		paths = append(paths, findPathRec(nodes, start, ends, instructions, 0))
	}

	println("Part 2: ", calculateLeastCommonMultiple(paths))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func calculateLeastCommonMultiple(nums []int) int {
	if len(nums) == 2 {
		return LCM(nums[0], nums[1])
	} else {
		return LCM(nums[0], calculateLeastCommonMultiple(nums[1:]))
	}
}

func findPathRec(nodes map[string]Node, start string, ends []string, instructions string, countMove int) int {
	if slices.Contains(ends, start) {
		return countMove
	}
	direction := instructions[countMove%len(instructions)]
	if direction == 'L' {
		start = nodes[start].Left
	} else {
		start = nodes[start].Right
	}

	return findPathRec(nodes, start, ends, instructions, countMove+1)
}

func part1(lineReader []string) {
	instructions := lineReader[0]
	usefulLines := lineReader[2:]

	graph := Graph{Vertices: map[string]*Vertex{}}

	counterPair := []string{}
	pairs := make([][]string, 0)
	keys := make([]string, 0)
	for _, line := range usefulLines {
		cleanedLine := utils.RemoveSpaces(line)
		cleanedLine = strings.Replace(cleanedLine, "(", "", -1)
		cleanedLine = strings.Replace(cleanedLine, ")", "", -1)
		splitLine := strings.Split(cleanedLine, "=")

		key := splitLine[0]
		pair := strings.Split(splitLine[1], ",")

		if key == "AAA" {
			counterPair = pair
		}
		keys = append(keys, key)

		pairs = append(pairs, pair)
		graph.AddVertex(key, pair)
	}

	counterKey := "AAA"

	println("Parsed vertexes: ", len(graph.Vertices))

	c := 1

Outer:
	for {
		for i := 0; i < len(instructions); i++ {
			inst := string(instructions[i])
			if inst == "L" {
				graph.AddEdge(counterKey, counterPair[0])
				counterKey = counterPair[0]
				counterPair = graph.Vertices[counterPair[0]].Val
			} else if inst == "R" {
				graph.AddEdge(counterKey, counterPair[1])
				counterKey = counterPair[1]
				counterPair = graph.Vertices[counterPair[1]].Val
			}

			if counterKey == "ZZZ" {
				break Outer
			}
			c++

		}

	}

	println("Part vertexes: ", len(graph.Vertices), " ", c)

}
