package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"dayeight/helpers"
)

type po struct {
	pattern []string
	output []string
}

func getInput(filename string) []po {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]po, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		patternOutput := strings.Split(scanner.Text(), " | ")
		input = append(input, po{
			strings.Split(patternOutput[0], " "),
			strings.Split(patternOutput[1], " "),
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func solution(input []po) int {
	return 0
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
	fmt.Println(helpers.GetAllDistances())
}
