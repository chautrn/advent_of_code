package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	//"strconv"
)

type stack struct {
	array []string
}

func (s *stack) add(c string) {
	s.array = append(s.array, c)
}

func (s *stack) next() string {
	if len(s.array) > 0 {
		return s.array[len(s.array) - 1]
	} else {
		return ""
	}
}

func (s *stack) pop() string {
	if len(s.array) > 0 {
		poppedElement := s.next()
		s.array = s.array[:len(s.array) - 1]
		return poppedElement
	}
	return ""
}

func isBigCave(cave string) bool {
	for _, c := range cave {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
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

func makeAdjs(input []string) map[string][]string {
	adjMap := make(map[string][]string)
	for _, relationship := range input {
		tokens := strings.Split(relationship, "-")
		for _, token := range tokens {
			if _, exists := adjMap[token]; !exists {
				adjList := make([]string, 0)
				adjMap[token] = adjList
			}
		}
		adjMap[tokens[0]] = append(adjMap[tokens[0]], tokens[1])
		adjMap[tokens[1]] = append(adjMap[tokens[1]], tokens[0])
	}
	return adjMap
}

func makeVisited(input []string) map[string]bool {
	visitedMap := make(map[string]bool)
	for _, relationship := range input {
		tokens := strings.Split(relationship, "-")
		for _, token := range tokens {
			visitedMap[token] = false
		}
	}
	return visitedMap
}

func traverse(cave string, pathBuffer stack, adjs map[string][]string, visited *map[string]bool, validPaths *[][]string, twice bool) {
	if !isBigCave(cave) {
		(*visited)[cave] = true
	}
	pathBuffer.add(cave)
	if cave == "end" {
		bufferCopy := append([]string(nil), pathBuffer.array...)
		*validPaths = append(*validPaths, bufferCopy)
	} else {
		for _, adj := range adjs[cave] {
			if !(*visited)[adj] {
				traverse(adj, pathBuffer, adjs, visited, validPaths, twice)
				if !twice && cave != "start" && cave != "end" && !isBigCave(cave) {
					(*visited)[cave] = false
					traverse(adj, pathBuffer, adjs, visited, validPaths, true)
					(*visited)[cave] = true
				}
			}
		}
	}
	if cave != "start"{
		(*visited)[cave] = false
	}
}

func listToString(input []string) string {
	result := ""
	for _, s := range input {
		result += s
	}
	return result
}

func removeDuplicates(input [][]string) []string {
	appeared := make(map[string]int)
	for _, l := range input {
		s := listToString(l)
		appeared[s] = 0
	}
	uniques := make([]string, 0)
	for k := range appeared {
		uniques = append(uniques, k)
	}
	return uniques
}

func solution(input []string) int {
	pathBuffer := stack{}
	adjMap := makeAdjs(input)
	visitedMap := makeVisited(input)
	validPaths := make([][]string, 0)
	traverse("start", pathBuffer, adjMap, &visitedMap, &validPaths, false)
	uniqueValidPaths := removeDuplicates(validPaths)
	return len(uniqueValidPaths)
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
