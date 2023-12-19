package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/STRockefeller/go-mine-sweeper/internal/sweeper"
	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
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
		color.Red(err.Error())
		os.Exit(1)
	}

	fmt.Printf("there are %d mines in the board.\r\n", board.Mines)

	scanner := bufio.NewScanner(os.Stdin)

	color.BlueString("hello world")
	previousX := -1
	previousY := -1
	// Game loop
	for !board.Display(previousX, previousY) {
		color.Cyan("Enter your move (e.g., 'R 3 4' to reveal or 'F 2 1' to flag) (0-indexed from the top-left corner): ")
		scanner.Scan()
		input := scanner.Text()

		commands := strings.Split(input, " ")
		if len(commands) != 3 {
			color.Red("invalid input")
			continue
		}

		x, err := strconv.Atoi(commands[1])
		if err != nil {
			color.Red("invalid input")
			continue
		}
		y, err := strconv.Atoi(commands[2])
		if err != nil {
			color.Red("invalid input")
			continue
		}

		if strings.ToUpper(commands[0]) == "R" {
			previousX = x
			previousY = y
			board.RevealCell(x, y)
		}

		if strings.ToUpper(commands[0]) == "F" {
			previousX = x
			previousY = y
			board.FlagCell(x, y)
		}
	}
}
