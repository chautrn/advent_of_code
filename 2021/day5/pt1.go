package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type position struct{
	x int
	y int
}

type line struct{
	a position
	b position
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func makePosition(sp string) position {
	stringComponents := strings.Split(sp, ",")
	x, err := strconv.Atoi(stringComponents[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(stringComponents[1])
	if err != nil {
		log.Fatal(err)
	}
	return position{x, y}
}

func makeLine(s string) line {
	stringPositions := strings.Split(s, " -> ")
	a := makePosition(stringPositions[0])
	b := makePosition(stringPositions[1])
	return line{a, b}
}

func getInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func formatInput(input []string) []line {
	formattedInput := make([]line, 0)
	for _, s := range input {
		line := makeLine(s)
		formattedInput = append(formattedInput, line)
	}
	return formattedInput
}

func isStraightLine(l line) bool {
	return l.a.x == l.b.x || l.a.y == l.b.y
}

func filterStraightLines(input []line)[]line {
	filteredInput := make([]line, 0)
	for _, l := range input {
		if isStraightLine(l) {
			filteredInput = append(filteredInput, l)
		}
	}
	return filteredInput
}

func drawVLine(x, startY, endY int, grid *[1000][1000]int) {
	for y := startY; y < endY + 1; y++ {
		grid[y][x] += 1
	}
}

func drawHLine(y, startX, endX int, grid *[1000][1000]int) {
	for x:= startX; x < endX + 1; x++ {
		grid[y][x] += 1
	}
}

func solution(input []string) int {
	formattedInput := formatInput(input)
	filteredInput := filterStraightLines(formattedInput)

	var grid [1000][1000]int

	for _, l := range filteredInput {
		// vertical line
		if l.a.x == l.b.x {
			start := min(l.a.y, l.b.y)
			end := max(l.a.y, l.b.y)
			drawVLine(l.a.x, start, end, &grid)
		// horizontal line
		} else if l.a.y == l.b.y {
			start := min(l.a.x, l.b.x)
			end := max(l.a.x, l.b.x)
			drawHLine(l.a.y, start, end, &grid)
		}
	}

	sum := 0

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if grid[y][x] >= 2 {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
