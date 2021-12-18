package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

const Max = 4294967295

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func getInput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()
	tokens := strings.Split(input, ",")
	parsedTokens := make([]int, 0)
	for _, token := range tokens {
		parsedToken, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		parsedTokens = append(parsedTokens, parsedToken)
	}
	return parsedTokens
}

func getSumLoss(x int, data []int) int {
	sumLoss := 0
	for _, p := range data {
		sumLoss += abs(p - x)
	}
	return sumLoss
}

func solution(input []int) int {
	min := Max
	for i := 0; i < 2000; i++ {
		fuel := getSumLoss(i, input)
		if fuel < min {
			min = fuel
		}
	}
	return min
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
