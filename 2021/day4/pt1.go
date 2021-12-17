package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type square struct {
	value	int
	marked	bool	`default: false`
}

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

func getDraw(line string)[]int {
	row := make([]int, 0)
	tokens := strings.Split(line, ",")
	for _, i := range tokens {
		parsedToken, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		row = append(row, parsedToken)
	}
	return row
}

func getRow(line string)[]square {
	row := make([]square, 0)
	tokens := strings.Split(line, " ")
	for _, i := range tokens {
		parsedToken, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		row = append(row, square{value: parsedToken})
	}
	return row
}

func getBoard(boardInput []string)[][]square {
	board := make([][]square, 0)
	for _, i := range boardInput {
		board = append(board, getRow(i))
	}
	return board
}

func getAllBoards(input []string)[][][]square {
	boards := make([][][]square, 0)
	for i := 2; i < len(input); i += 6 {
		boards = append(boards, getBoard(input[i:i+5]))
	}
	return boards
}

func processSingleDraw(draw int, board *[][]square) {
	for i, _ := range *board {
		for j, _ := range (*board)[i] {
			if (*board)[i][j].value == draw {
				(*board)[i][j].marked = true
			}
		}
	}
}

func evaluateBoard(board *[][]square)bool {
	for ci, _ := range *board {
		if (*board)[0][ci].marked == true {
			rowMatch := true
			for ri, _ := range (*board) {
				if !((*board)[ri][ci].marked) {
					rowMatch = false
				}
			}
			if rowMatch {
				return rowMatch
			}
		}
	}

	for ri, _ := range *board {
		if (*board)[ri][0].marked == true {
			colMatch := true
			for ci, _ := range (*board)[ri] {
				if !((*board)[ri][ci].marked) {
					colMatch = false
				}
			}
			if colMatch {
				return colMatch
			}
		}
	}

	return false
}

func sumUnmarked(board *[][]square)int {
	sum := 0
	for _, row := range *board {
		for _, element := range row {
			if !element.marked {
				sum += element.value
			}
		}
	}

	return sum
}

func solution(input []string)int {
	draws := getDraw(input[0])
	boards := getAllBoards(input)

	for i, draw := range draws {
		for _, board := range boards {
			processSingleDraw(draw, &board)
			if evaluateBoard(&board) {
				return draw * sumUnmarked(&board)
			}
		}
	}

	return 0
}

func main() {
	input := getInput("./input.txt")
	fmt.Println(solution(input))
}
