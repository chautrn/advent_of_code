package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type queue struct {
	slice []int
}

func (q *queue) pop() int {
	prev := q.slice[0]
	var temp int
	for i := 1; i < len(q.slice); i++ {
		temp = q.slice[i]
		q.slice[i] = prev
		prev = temp
	}
	q.slice[0] = 0
	return prev
}

func sum(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
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

func solution(input []int, days int) int {
	// indices: 	[0 1 2 3 4 5 6 7 8]
	// fish days:	[8 7 6 5 4 3 2 1 0]
	fish := make([]int, 9)
	for _, i := range input {
		fish[8 - i] += 1
	}
	birthCycle := queue{fish}
	for i := 0; i < days; i++ {
		fishToBeResetted := birthCycle.pop()
		birthCycle.slice[2] += fishToBeResetted
		birthCycle.slice[0] += fishToBeResetted
	}
	return sum(birthCycle.slice)
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input, 256))
}
