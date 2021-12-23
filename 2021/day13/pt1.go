package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type point struct {
	x int
	y int
}

func getInput(filename string) ([]string, []string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := make([]string, 0)
	folds := make([]string, 0)
	p := &points
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			p = &folds
			continue
		}
		*p = append(*p, text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points, folds
}

func parsePoints(pointStrings []string) []point {
	points := make([]point, 0)
	for _, str := range pointStrings {
		tokens := strings.Split(str, ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, point{x, y})
	}
	return points
}

func solution(input [][]int) int {
	return 0
}

func main() {
	pointStrings, foldStrings := getInput("./input.txt")
	points := parsePoints(pointStrings)
	fmt.Println(points)
	fmt.Println(foldStrings)
}
