package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput(filename string)[]string {
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

func getBit(count, threshold int, opposite bool)byte {
	isPositive := count >= threshold
	if opposite {
		isPositive = !isPositive
	}
	if isPositive {
		return '1'
	} else {
		return '0'
	}
}

func filter(input []string, index int, opposite bool)[]string {
	result := make([]string, 0)
	count := 0
	for _, s := range input {
		if s[index] == '1' {
			count++
		}
	}
	bit := getBit(count, (len(input) + 1) / 2, opposite)
	for _, s := range input {
		if s[index] == bit {
			result = append(result, s)
		}
	}
	return result
}

func applyFilter(input []string, isCO2 bool)string {
	result := input
	for i := 0; i < len(input[0]); i++ {
		result = filter(result, i, isCO2)
		if len(result) == 1 {
			return result[0]
		}
	}
	return result[0]
}

func product(input []string)int {
	ogRating, err := strconv.ParseUint(applyFilter(input, false), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	csRating, err := strconv.ParseUint(applyFilter(input, true), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(ogRating * csRating)
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(product(input))
}
