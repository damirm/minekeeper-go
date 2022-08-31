package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	HEIGHT = 15
	WIDTH  = 15
	// Bomb rate percents
	BOMB_RATE = 20
)

type CellState int

const (
	CellStateClosed CellState = iota
	CellStateOpened
	CellStateMarked
	CellStateUnknown
)

type Point struct {
	x, y int
}

type Board2 struct {
	gameover bool
	cells    [][]int
	states   [][]CellState
	width    int
	height   int
	cursor   *Point
}

func NewBoard2(width, height int) *Board2 {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
	}
	states := make([][]CellState, width)
	for i := range states {
		states[i] = make([]CellState, width)
	}
	return &Board2{
		cells:  cells,
		states: states,
		width:  width,
		height: height,
		cursor: &Point{0, 0},
	}
}

func (b *Board2) OpenCell(x, y int) bool {
	if b.cells[y][x] == 1 {
		b.gameover = true
	}
	b.states[y][x] = CellStateOpened
	return b.gameover
}

func (b *Board2) GetWidth() int {
	return b.width
}

func (b *Board2) GetHeight() int {
	return b.height
}

func (b *Board2) GetCellState(x, y int) CellState {
	return b.states[y][x]
}

func (b *Board2) Randomize(bombRate float64) {
	bombsCount := math.Floor(float64(b.height*b.width) * (bombRate / 100.0))
	cnt := 0.0
	for cnt < bombsCount {
		x, y := random(0, WIDTH), random(0, HEIGHT)
		if b.cells[y][x] == 1 {
			continue
		}

		b.cells[y][x] = 1
		cnt++
	}
}

func (b *Board2) IsCursorAt(x, y int) bool {
	return b.cursor.x == x && b.cursor.y == y
}

func (b *Board2) IsBombAt(x, y int) bool {
	return b.cells[y][x] == 1
}

func (b *Board2) MoveCursor(dx, dy int) {
	x := b.cursor.x + dx
	y := b.cursor.y + dy

	if x < 0 {
		x = 0
	}
	if x >= b.width {
		x = b.width - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= b.height {
		y = b.height - 1
	}

	b.cursor.x = x
	b.cursor.y = y
}

func (b *Board2) CountBombsAround(x, y int) int {
	var res int

	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if iy != 0 && ix != 0 {
				nx := x + ix
				ny := y + iy

				if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height {
					if b.cells[ny][nx] == 1 {
						res++
					}
				}
			}
		}
	}

	return res
}

type Board [HEIGHT][WIDTH]int

var cellState [HEIGHT][WIDTH]CellState

// cursor position
var cx, cy int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func openCell(board *Board, i, j int) bool {
	cellState[i][j] = CellStateOpened
	return board[i][j] == 1
}

func random(from, to int) int {
	return rand.Intn(to) + from
}

func putRandomBombs(board *Board) {
	bombsCount := math.Floor(HEIGHT * WIDTH * (BOMB_RATE / 100.0))
	cnt := 0.0
	for cnt < bombsCount {
		x, y := random(0, WIDTH), random(0, HEIGHT)
		if board[y][x] == 1 {
			continue
		}

		board[y][x] = 1
		cnt++
	}
}

// 0 0 0
// 0 0 0
// 0 0 0
func countBombsAround(board *Board, y, x int) int {
	var res int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			ny := y + i
			nx := x + j

			if i != 0 && j != 0 && nx >= 0 && nx < WIDTH && ny >= 0 && ny < HEIGHT {
				if board[ny][nx] == 1 {
					res++
				}
			}
		}
	}
	return res
}

func printBoard(board *Board) {
	fmt.Print("\u001B[G")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			var cell string
			switch cellState[i][j] {
			case CellStateClosed:
				cell = "*"
			case CellStateOpened:
				// case CellStateClosed:
				if board[i][j] == 1 {
					cell = "@"
				} else {
					cell = fmt.Sprintf("%d", countBombsAround(board, i, j))
				}
			}

			if i == cy && j == cx {
				cell = "[" + cell + "]"
			} else {
				cell = " " + cell + " "
			}

			fmt.Print(cell)
		}
		fmt.Print("\u001B[G\n")
	}
	fmt.Printf("\u001B[%dD\u001B[%dA", WIDTH, HEIGHT)
}

func loop() {
	rand.Seed(time.Now().UnixNano())

	var board Board
	putRandomBombs(&board)
	printBoard(&board)

	var reader = bufio.NewReaderSize(os.Stdin, 1)
	var quit = false

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case 'h':
			cx = max(0, cx-1)
		case 'j':
			cy = min(HEIGHT-1, cy+1)
		case 'k':
			cy = max(0, cy-1)
		case 'l':
			cx = min(WIDTH-1, cx+1)
		case 'q':
			quit = true
		case ' ':
			quit = openCell(&board, cy, cx)
		}

		if quit {
			break
		}

		printBoard(&board)
	}
}

func main() {
	state, err := term.MakeRaw(0)
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(0, state)

	loop()
}
