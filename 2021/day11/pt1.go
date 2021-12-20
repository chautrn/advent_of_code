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
		intMatrix = append(intMatrix, rowParsed)
	}
	return intMatrix
}

func inBounds(y, x int, input *[][]int) bool {
	yInBounds := y >= 0 && y < len((*input))
	xInBounds := x >= 0 && x < len((*input)[0])
	return yInBounds && xInBounds
}

func propagateFlash(y, x int, input *[][]int) int {
	sum := 1
	(*input)[y][x] = 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if inBounds(y + j, x + i, input) {
				if (*input)[y + j][x + i] != 0 {
					(*input)[y + j][x + i] += 1
					if (*input)[y + j][x + i] > 9 {
						sum += propagateFlash(y + j, x + i, input)
					}
				}
			}
		}
	}
	return sum
}

func executeFlashes(input *[][]int) int {
	sum := 0
	for j := range *input {
		for i := range (*input)[0] {
			if (*input)[j][i] > 9 {
				sum += propagateFlash(j, i, input)
			}
		}
	}
	return sum
}

func increaseEnergy(input *[][]int) {
	for j := range *input {
		for i := range (*input)[0] {
			(*input)[j][i] += 1
		}
	}
}

func printGrid(input [][]int) {
	for j := range input {
		for i := range input[0] {
			fmt.Printf("%x", input[j][i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func solution(input [][]int, steps int) int {
	sum := 0
	for i := 0; i < steps; i++ {
		fmt.Printf("step %d\n", i)
		printGrid(input)
		increaseEnergy(&input)
		sum += executeFlashes(&input)
	}
	return sum
}

func main() {
	input := getInput("./input.txt")
	formattedInput := formatInput(input)
	fmt.Println(solution(formattedInput, 100))
}
