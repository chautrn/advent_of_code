package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"strconv"
)

func getInput(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		afterDel := strings.Split(scanner.Text(), " | ")[1]
		tokens := strings.Split(afterDel, " ")
		input = append(input, tokens)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func solution(input [][]string) int {
	sum := 0
	for _, segment := range input {
		for _, digit := range segment {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
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
