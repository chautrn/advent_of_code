package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

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

func nList(element, numElements int) []int {
	result := make([]int, numElements)
	for i, _ := range result {
		result[i] = element
	}
	return result
}

func formatInput(input []string) [][]int {
	intMatrix := make([][]int, 0)
	for _, row := range input {
		rowData := strings.Split(row, "")
		rowParsed := make([]int, len(rowData))
		for i, numString := range rowData {
			var numParsed int
			numParsed, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			} else {
				rowParsed[i] = numParsed
			}
		}
		rowParsed = append([]int{9}, rowParsed...) // left padding
		rowParsed = append(rowParsed, 9) // right padding
		intMatrix = append(intMatrix, rowParsed)
	}
	padding := nList(9, len(intMatrix[0]))
	intMatrix = append([][]int{padding}, intMatrix...)
	intMatrix = append(intMatrix, padding)
	return intMatrix
}

func lowKernel(y, x int, input[][]int) bool {
	yAxisLow := input[y][x] < input[y - 1][x] && input[y][x] < input[y + 1][x]
	xAxisLow := input[y][x] < input[y][x - 1] && input[y][x] < input[y][x + 1]
	return yAxisLow && xAxisLow
}

func inBounds(y, x int, input *[][]int) bool {
	yInBounds := y >= 0 && y < len(*input)
	xInBounds := x >= 0 && x < len((*input)[0])
	return yInBounds && xInBounds
}

func notBoundary(y, x int, input *[][]int) bool {
	return (*input)[y][x] != -1 && (*input)[y][x] != 9
}

func validPoint(y, x int, input *[][]int) bool {
	return inBounds(y, x, input) && notBoundary(y, x, input)
}

func growRegion(y, x int, input *[][]int) int {
	inputMatrix := *input
	inputMatrix[y][x] = -1
	sum := 1

	// grow y-direction
	if validPoint(y + 1, x, input) {
		inputMatrix[y + 1][x] = -1
		sum += growRegion(y + 1, x, input)
	}
	if validPoint(y - 1, x, input) {
		inputMatrix[y - 1][x] = -1
		sum += growRegion(y - 1, x, input)
	}

	// grow x-direction
	if validPoint(y, x + 1, input) {
		inputMatrix[y][x + 1] = -1
		sum += growRegion(y, x + 1, input)
	}
	if validPoint(y, x - 1, input) {
		inputMatrix[y][x - 1] = -1
		sum += growRegion(y, x - 1, input)
	}

	return sum
}

func solution(input [][]int) int {
	topSizes := make([]int, 0)
	for j := 1; j < len(input) - 1; j++ {
		for i := 1; i < len(input[0]) - 1; i++ {
			if lowKernel(j, i, input) {
				size := growRegion(j, i, &input)
				topSizes = append(topSizes, size)
			}
		}
	}
	sort.Ints(topSizes)
	l := len(topSizes)
	return topSizes[l - 1] * topSizes[l - 2] * topSizes[l - 3]
}

func main() {
	input := getInput("./input.txt")
	formattedInput := formatInput(input)
	fmt.Println(solution(formattedInput))
}
