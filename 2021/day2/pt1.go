package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	a *int
	b int
}

func getProduct(course []string) int {
	x := 0
	y := 0

	o := make(map[string] pair)
	o["up"] = pair{&y, -1}
	o["down"] = pair{&y, 1}
	o["forward"] = pair{&x, 1}

	for _, el := range course {
		tokens := strings.Split(el, " ")
		direction := tokens[0]
		scalar, err := strconv.Atoi(tokens[1])

		if err != nil {
			log.Fatal(err)
			continue
		}

		*o[direction].a += scalar * o[direction].b
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
