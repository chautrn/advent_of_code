package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"strconv"
)

type stack struct {
	array []string
}

func (s *stack) add(c string) {
	s.array = append(s.array, c)
}

func (s *stack) pop() string {
	if len(s.array) > 0 {
		poppedElement := s.array[len(s.array) - 1]
		s.array = s.array[:len(s.array) - 1]
		return poppedElement
	}
	return ""
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

func isOpen(c string) bool {
	return strings.Contains("([{<", c)
}

func compliment(c string) string {
	switch c {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	return ""
}

func firstIllegalCharacter(s string) string {
	expecting := stack{}
	for _, c := range s {
		char := string(c)
		if isOpen(char) {
			expecting.add(compliment(char))
		} else {
			if expecting.pop() != char {
				return char
			}
		}
	}
	return ""
}

func calcScore(c string) int {
	switch c {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func solution(input []string) int {
	score := 0
	for _, line := range input {
		char := firstIllegalCharacter(line)
		score += calcScore(char)
	}
	return score
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
