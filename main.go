package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type state int

const (
	stateOff state = 0
	stateOn  state = 1

	cellSize = 25
)

var board [][]state
var newBoard [][]state
var win *pixelgl.Window
var tile *imdraw.IMDraw

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	board = readInput()
	newBoard = make([][]state, len(board))
	for row := 0; row < len(board); row++ {
		newBoard[row] = append([]state(nil), board[row]...)
	}

	pixelgl.Run(run)
}

func readInput() [][]state {
	// return [][]state{
	// 	{stateOn, stateOff, stateOn, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOff, stateOn, stateOff},
	// 	{stateOn, stateOff, stateOn, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOff, stateOn, stateOff},
	// }

	// return [][]state{
	// 	{stateOn, stateOff, stateOn, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOff},
	// }

	// Blinker
	// return [][]state{
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff},
	// }

	// Toad
	// return [][]state{
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOn, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOn, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// }

	// Beacon
	// return [][]state{
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOn, stateOn, stateOff, stateOff, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOn, stateOn, stateOff},
	// 	{stateOff, stateOff, stateOff, stateOff, stateOff, stateOff},
	// }

	// Pulsar
	// return [][]state{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }

	// Random
	b := make([][]state, 30)
	for row := 0; row < 30; row++ {
		b[row] = make([]state, 30)
		for col := 0; col < 30; col++ {
			b[row][col] = state(rand.Intn(2))
		}
	}

	return b
}

func run() {
	initPixel()
	drawAll()
	startLoop()
}

// Initialize pixel, sprite, etc
func initPixel() {
	boardHeight := len(board)
	boardWidth := len(board[0])
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, float64(boardWidth*cellSize), float64(boardHeight*cellSize)),
		VSync:  true,
	}

	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tile = imdraw.New(nil)
}

// Will start the handling of events loop
func startLoop() {
	last := time.Now()

	for !win.Closed() {
		dtMilliS := time.Since(last).Milliseconds()
		if dtMilliS > 1000 { // Every second, shift the state
			updateCells()
			draw()
			for row := 0; row < len(board); row++ {
				board[row] = append([]state(nil), newBoard[row]...)
			}
			last = time.Now()
		}
		win.Update()
	}

}

// updateCells updates all cells
func updateCells() {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			updateCell(row, col)
		}
	}
}

// updateCell updates the given cell
func updateCell(row, col int) {
	nb := countNeighbours(row, col)
	if (board[row][col] == stateOn && (nb == 2 || nb == 3)) ||
		(board[row][col] == stateOff && nb == 3) {
		newBoard[row][col] = stateOn
	} else {
		newBoard[row][col] = stateOff
	}
}

// countNeighbours will count the number of neighbours of the given cell
func countNeighbours(row, col int) int {
	return cellValue(row-1, col-1) + cellValue(row-1, col) + cellValue(row-1, col+1) +
		cellValue(row, col-1) + cellValue(row, col+1) +
		cellValue(row+1, col-1) + cellValue(row+1, col) + cellValue(row+1, col+1)
}

// cellValue will return value of the cell
func cellValue(row, col int) int {
	if row >= 0 &&
		row < len(board) &&
		col >= 0 &&
		col < len(board[0]) &&
		board[row][col] == stateOn {
		return 1
	}

	return 0
}

// drawAll will draw all cells
func drawAll() {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			drawCell(row, col)
		}
	}
}

// draw draws only the updated cells
func draw() {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if newBoard[row][col] != board[row][col] {
				drawCell(row, col)
			}
		}
	}
}

// drawCell is to draw a single cell
func drawCell(row, col int) {
	color := colornames.Red
	if newBoard[row][col] == stateOff {
		color = colornames.Beige
	}

	tile.Color = color
	tile.Push(pixel.V(float64(col)*cellSize+3, float64(row)*cellSize+3))
	tile.Push(pixel.V(float64(col+1)*cellSize-3, float64(row+1)*cellSize-3))
	tile.Rectangle(0)
	tile.Draw(win)
}
