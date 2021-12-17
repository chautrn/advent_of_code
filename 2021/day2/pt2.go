package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getProduct(course []string) int {
	x := 0
	y := 0
	aim := 0

	for _, el := range course {
		tokens := strings.Split(el, " ")
		direction := tokens[0]
		scalar, err := strconv.Atoi(tokens[1])

		if err != nil {
			log.Fatal(err)
			continue
		}

		switch direction {
			case "forward":
				x += scalar
				y += aim * scalar
			case "up":
				aim -= scalar
			case "down":
				aim += scalar
		}
	}

	return x * y
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

	product := getProduct(input)
	fmt.Printf("%d\n", product)
}
