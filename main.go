package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	HasMine          bool
	Revealed         bool
	HasFlag          bool
	NeighboringMines int
}

type Board struct {
	Grid          [][]Cell
	Width, Height int
	Mines         int
}

func NewBoard(width, height, mines int) *Board {
	board := &Board{
		Grid:   make([][]Cell, height),
		Width:  width,
		Height: height,
		Mines:  mines,
	}

	for i := range board.Grid {
		board.Grid[i] = make([]Cell, width)
	}

	// Add mines at random positions
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for m := 0; m < mines; m++ {
		x, y := rand.Intn(width), rand.Intn(height)
		// Check if there's already a mine here
		if board.Grid[y][x].HasMine {
			m-- // Retry
			continue
		}
		board.Grid[y][x].HasMine = true
	}

	return board
}

func (b *Board) calculateNeighboringMines() {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Grid[y][x].HasMine {
				continue
			}
			for i := 0; i < 8; i++ {
				nx, ny := x+dx[i], y+dy[i]
				if nx >= 0 && ny >= 0 && nx < b.Width && ny < b.Height && b.Grid[ny][nx].HasMine {
					b.Grid[y][x].NeighboringMines++
				}
			}
		}
	}
}

func (b *Board) Display() (gameOver bool) {
	unrevealedGrids := 0
	endGameMessage := "YOU WON THE GAME!"
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			cell := b.Grid[y][x]
			if cell.Revealed {
				if cell.HasMine {
					fmt.Print("* ")
					endGameMessage = "BOMB!"
					gameOver = true
				} else {
					fmt.Printf("%d ", cell.NeighboringMines)
				}
			} else if cell.HasFlag {
				fmt.Print("F ")
				unrevealedGrids++
			} else {
				fmt.Print(". ")
				unrevealedGrids++
			}
		}
		fmt.Println()
	}
	if unrevealedGrids == b.Mines {
		gameOver = true
	}
	if gameOver {
		fmt.Println(endGameMessage)
	}
	return
}

func (b *Board) RevealCell(x, y int) {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return
	}

	cell := &b.Grid[y][x]
	if cell.Revealed {
		return
	}

	cell.Revealed = true
	if cell.HasMine {
		return
	}

	if cell.NeighboringMines == 0 {
		// Reveal neighboring cells
		dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
		dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}
		for i := 0; i < 8; i++ {
			b.RevealCell(x+dx[i], y+dy[i])
		}
	}
}

func (b *Board) FlagCell(x, y int) {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return
	}
	b.Grid[y][x].HasFlag = !b.Grid[y][x].HasFlag
}

func main() {
	width, height, mines := 10, 10, 10
	board := NewBoard(width, height, mines)
	board.calculateNeighboringMines()

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
