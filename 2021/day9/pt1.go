package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
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

func solution(input [][]int) int {
	sum := 0
	for j := 1; j < len(input) - 1; j++ {
		for i := 1; i < len(input[0]) - 1; i++ {
			if lowKernel(j, i, input) {
				sum += input[j][i] + 1
			}
		}
	}
	return sum
}

func main() {
	input := getInput("./input.txt")
	formattedInput := formatInput(input)
	fmt.Println(solution(formattedInput))
}
