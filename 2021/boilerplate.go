package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
	//"strconv"
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

func solution(input []string) int {
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
