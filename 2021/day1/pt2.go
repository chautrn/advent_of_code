package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func stringArrayToInt(array []string) []int {
	intArray := make([]int, 0)
	for _, s := range array {
		if intValue, err := strconv.Atoi(s); err == nil {
			intArray = append(intArray, intValue)
		}
	}
	return intArray
}

func calcWindow(index int, slice []int) int {
	return slice[index - 1] + slice[index] + slice[index + 1]
}

func countIncreases(depths []int) int {
	increases := 0

	for i := 2; i < len(depths) - 1; i++ {
		previous := calcWindow(i - 1, depths)
		current := calcWindow(i, depths)
		if current > previous {
			increases++
		}
	}

	return increases
}

func main() {
	f, err := os.Open("./input.txt")
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

	intInput := stringArrayToInt(input)
	increases := countIncreases(intInput)
	fmt.Printf("%d\n", increases)
}
