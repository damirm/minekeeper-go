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

type Board struct {
	gameover bool
	cells    [][]int
	states   [][]CellState
	width    int
	height   int
	cursor   *Point
}

func NewBoard(width, height int) *Board {
	b := &Board{
		width:  width,
		height: height,
		cursor: &Point{0, 0},
	}
	b.init()
	return b
}

func (b *Board) init() {
	b.cells = make([][]int, b.height)
	for i := range b.cells {
		b.cells[i] = make([]int, b.width)
	}
	b.states = make([][]CellState, b.width)
	for i := range b.states {
		b.states[i] = make([]CellState, b.width)
	}
}

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) GetCellState(x, y int) CellState {
	return b.states[y][x]
}

func (b *Board) Randomize(bombRate float64) {
	b.init()

	bombsCount := math.Floor(float64(b.height*b.width) * (bombRate / 100.0))
	cnt := 0.0
	for cnt < bombsCount {
		x, y := random(0, WIDTH), random(0, HEIGHT)
		if b.cells[y][x] == 1 {
			continue
		}

		b.PutBombTo(x, y)
		cnt++
	}
}

func (b *Board) PutBombTo(x, y int) {
	b.cells[y][x] = 1
}

func (b *Board) ResetCells() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.cells[i][j] = 0
		}
	}
}

func (b *Board) OpenCurrentCell() bool {
	b.openCell(b.cursor.x, b.cursor.y)
	if !b.gameover {
		b.expandNeighbors(b.cursor.x, b.cursor.y)
	}
	return b.gameover
}

func (b *Board) openCell(x, y int) {
	if b.cells[y][x] == 1 {
		b.gameover = true
	}
	b.states[y][x] = CellStateOpened
}

func (b *Board) expandNeighbors(x, y int) {
	if b.CountBombsAround(x, y) == 0 {
		b.loopAround(x, y, func(nx, ny int) {
			if b.states[ny][nx] != CellStateOpened {
				b.openCell(nx, ny)
				b.expandNeighbors(nx, ny)
			}
		})
	}
}

func (b *Board) OpenAllCells() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.states[i][j] = CellStateOpened
		}
	}
}

func (b *Board) MarkCurrentCell() {
	state := b.states[b.cursor.y][b.cursor.x]
	var nstate CellState
	switch state {
	case CellStateClosed:
		nstate = CellStateMarked
	case CellStateMarked:
		nstate = CellStateUnknown
	case CellStateUnknown:
		nstate = CellStateClosed
	default:
		nstate = CellStateClosed
	}
	b.states[b.cursor.y][b.cursor.x] = nstate
}

func (b *Board) IsCursorAt(x, y int) bool {
	return b.cursor.x == x && b.cursor.y == y
}

func (b *Board) IsBombAt(x, y int) bool {
	return b.cells[y][x] == 1
}

func (b *Board) MoveCursor(dx, dy int) {
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

func (b *Board) SetCursor(cursor *Point) {
	b.cursor = cursor
}

func (b *Board) CountBombsAround(x, y int) int {
	var res int

	b.loopAround(x, y, func(nx, ny int) {
		if b.cells[ny][nx] == 1 {
			res++
		}
	})

	return res
}

func (b *Board) loopAround(x, y int, cb func(int, int)) {
	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if iy != 0 || ix != 0 {
				nx := x + ix
				ny := y + iy

				if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height {
					cb(nx, ny)
				}
			}
		}
	}
}

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

func random(from, to int) int {
	return rand.Intn(to-from) + from
}

func printBoard(board *Board) {
	fmt.Print("\u001B[G")
	for i := 0; i < board.GetHeight(); i++ {
		for j := 0; j < board.GetWidth(); j++ {
			var cell string
			switch board.GetCellState(j, i) {
			case CellStateClosed:
				cell = "."
			case CellStateOpened:
				if board.IsBombAt(j, i) {
					cell = "@"
				} else {
					around := board.CountBombsAround(j, i)
					if around > 0 {
						cell = fmt.Sprintf("%d", around)
					} else {
						cell = " "
					}
				}
			case CellStateMarked:
				cell = "!"
			case CellStateUnknown:
				cell = "?"
			}

			if board.IsCursorAt(j, i) {
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

	board := NewBoard(WIDTH, HEIGHT)
	board.Randomize(BOMB_RATE)

	printBoard(board)

	var reader = bufio.NewReaderSize(os.Stdin, 1)
	var quit = false

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case 'h':
			board.MoveCursor(-1, 0)
		case 'j':
			board.MoveCursor(0, 1)
		case 'k':
			board.MoveCursor(0, -1)
		case 'l':
			board.MoveCursor(1, 0)
		case 'm':
			board.MarkCurrentCell()
		case 'q':
			quit = true
		case 'r':
			board.Randomize(BOMB_RATE)
		case ' ':
			quit = board.OpenCurrentCell()
		}

		if quit {
			break
		}

		printBoard(board)
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
