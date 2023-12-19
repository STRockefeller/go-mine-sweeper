package sweeper

import (
	"fmt"

	"github.com/fatih/color"

	"math/rand"
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

	board.calculateNeighboringMines()

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

func (b *Board) Display(highlightX, highlightY int) (gameOver bool) {
	unrevealedGrids := 0
	endGameMessage := "YOU WON THE GAME!"

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if x == highlightX && y == highlightY {
				color.Unset()
				color.Set(color.FgMagenta)
			}
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
			color.Unset()
		}
		fmt.Println()
	}
	if unrevealedGrids == b.Mines {
		gameOver = true
	}
	if gameOver {
		color.Red(endGameMessage)
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
