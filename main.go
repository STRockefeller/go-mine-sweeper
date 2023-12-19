package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/STRockefeller/go-mine-sweeper/internal/sweeper"
)

func main() {
	width, height, mines := 10, 10, 10
	board := sweeper.NewBoard(width, height, mines)

	scanner := bufio.NewScanner(os.Stdin)
	// Game loop
	for !board.Display() {
		fmt.Print("Enter your move (e.g., 'R 3 4' to reveal or 'F 2 1' to flag): ")
		scanner.Scan()
		input := scanner.Text()

		commands := strings.Split(input, " ")
		if len(commands) != 3 {
			fmt.Println("invalid input")
			continue
		}

		x, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		y, err := strconv.Atoi(commands[2])
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		if commands[0] == "R" {
			board.RevealCell(x, y)
		}

		if commands[0] == "F" {
			board.FlagCell(x, y)
		}
	}
}
