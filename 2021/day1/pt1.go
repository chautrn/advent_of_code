package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncreases(depths []string) int {
	increases := 0

	for i := 1; i < len(depths); i++ {
		previous, err := strconv.ParseInt(depths[i-1], 10, 64)
		current, err := strconv.ParseInt(depths[i], 10, 64)
		if err != nil {
			log.Fatal(err)
			continue
		} else if current > previous {
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

	increases := countIncreases(input)
	fmt.Printf("%d\n", increases)
}
