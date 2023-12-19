package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/STRockefeller/go-mine-sweeper/internal/sweeper"
	"github.com/charmbracelet/huh"
)

func main() {
	var board *sweeper.Board
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[*sweeper.Board]().Title("Game Difficulty").Options(
				huh.NewOption("Easy", sweeper.NewBoard(8, 8, 7)),
				huh.NewOption("Medium", sweeper.NewBoard(10, 10, 15)),
				huh.NewOption("Hard", sweeper.NewBoard(12, 12, 29)),
			).Value(&board),
		),
	)
	if err := form.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("there are %d mines in the board.\r\n", board.Mines)

	scanner := bufio.NewScanner(os.Stdin)
	// Game loop
	for !board.Display() {
		fmt.Print("Enter your move (e.g., 'R 3 4' to reveal or 'F 2 1' to flag) (0-indexed from the top-left corner): ")
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
