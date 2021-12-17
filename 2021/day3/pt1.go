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

func commonBinary(input []string)(string, string) {
	count := make([]int, len(input[0]))
	for _, s := range input {
		for i, l := range s {
			if l == '1' {
				count[i] += 1
			}
		}
	}
	gamma := ""
	epsilon := ""
	for _, c := range count {
		if c > (len(input)/2) {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	return gamma, epsilon
}

func powerConsumption(gamma, epsilon string)int {
	g, err := strconv.ParseUint(gamma, 2, 16)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseUint(epsilon, 2, 16)
	if err != nil {
		log.Fatal(err)
	}
	return int(g*e)
}

func main() {
	input := getInput("./input.txt")
	gamma, epsilon := commonBinary(input)
	p := powerConsumption(gamma, epsilon)
	fmt.Println(p)
}
