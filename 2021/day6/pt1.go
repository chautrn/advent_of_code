package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

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

func updateFish(input *[]int) {
	numNewFish := 0
	for i := range *input {
		if (*input)[i] == 0 {
			numNewFish += 1
			(*input)[i] = 6
		} else {
			(*input)[i] -= 1
		}
	}
	for i := 0; i < numNewFish; i++ {
		*input = append(*input, 8)
	}
}

func solution(input []int, days int) int {
	for day := 1; day < days + 1; day++ {
		updateFish(&input)
	}
	return len(input)
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input, 80))
}
