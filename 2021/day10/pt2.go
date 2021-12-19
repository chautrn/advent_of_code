package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"
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

func filterInput(input []string) []string {
	filtered := make([]string, 0)
	for _, line := range input {
		if len(firstIllegalCharacter(line)) == 0 {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func autoComplete(s string) string {
	expecting := stack{}
	completionString := ""
	for _, c := range s {
		char := string(c)
		if isOpen(char) {
			expecting.add(compliment(char))
		} else if expecting.pop() != char {
			return char
		}
	}
	for i := len(expecting.array) - 1; i >= 0; i-- {
		completionString += expecting.array[i]
	}
	return completionString
}

func calcScore(s string) int {
	score := 0
	scoreMap := make(map[rune]int)
	scoreMap[')'] = 1
	scoreMap[']'] = 2
	scoreMap['}'] = 3
	scoreMap['>'] = 4
	for _, c := range s {
		score *= 5
		score += scoreMap[c]
	}
	return score
}

func solution(input []string) int {
	scores := make([]int, 0)
	for _, line := range input {
		completionString := autoComplete(line)
		scores = append(scores, calcScore(completionString))
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	input := getInput("./input.txt")
	filteredInput := filterInput(input)
	fmt.Println(solution(filteredInput))
}
