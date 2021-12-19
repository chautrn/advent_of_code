package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"day8/helpers"
	"sort"
	"reflect"
	"strconv"
)

type po struct {
	pattern []string
	output []string
}

func equalsIgnoreOrder(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	return reflect.DeepEqual(a, b)
}

func alphabetizeString(s string) string {
	charSlice := strings.Split(s, "")
	sort.Strings(charSlice)
	return strings.Join(charSlice, "")
}

func alphabetizeStringSlice(strings []string) []string {
	for i, s := range strings {
		strings[i] = alphabetizeString(s)
	}
	return strings
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
		pattern := strings.Split(patternOutput[0], " ")
		output := strings.Split(patternOutput[1], " ")
		input = append(input, po{
			alphabetizeStringSlice(pattern),
			alphabetizeStringSlice(output),
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func getIndex(element []int, in [][]int) int {
	for i, potentialMatch := range in {
		if equalsIgnoreOrder(element, potentialMatch) {
			return i
		}
	}
	return -1
}

func mapStrings(input po) map[string]int {
	mappedDistances := make(map[string]int)
	distances := helpers.GetAllDistances()
	toBeMapped := helpers.GetAllAlphaDistances(input.pattern)
	for i, _ := range distances {
		at := getIndex(distances[i], toBeMapped)
		mappedDistances[input.pattern[at]] = i
	}
	return mappedDistances
}

func solution(input []po) int {
	sum := 0
	for _, patternOutput := range input {
		mapped := mapStrings(patternOutput)
		numString := ""
		for _, str := range patternOutput.output {
			numString += strconv.Itoa(mapped[str])
		}
		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	return sum
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
