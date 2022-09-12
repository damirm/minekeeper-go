package main

import "testing"

func point(x, y int) Point {
	return Point{x: x, y: y}
}

func putAllBombs(board *Board, bombs []Point) {
	for _, bomb := range bombs {
		board.PutBombTo(bomb.x, bomb.y)
	}
}

func TestCountBombsAround(t *testing.T) {
	for _, tc := range []struct {
		board  *Board
		bombs  []Point
		cursor Point
		exp    int
	}{
		{
			NewBoard(3, 3),
			[]Point{},
			point(1, 1),
			0,
		},
		{
			NewBoard(3, 3),
			[]Point{
				point(0, 0),
			},
			point(1, 1),
			1,
		},
		{
			NewBoard(3, 3),
			[]Point{
				point(0, 0),
				point(1, 0),
			},
			point(1, 1),
			2,
		},
		{
			NewBoard(3, 3),
			[]Point{
				point(0, 0),
				point(1, 0),
				point(2, 2),
			},
			point(1, 1),
			3,
		},
	} {
		putAllBombs(tc.board, tc.bombs)

		tc.board.SetCursor(&tc.cursor)
		tc.board.OpenAllCells()

		got := tc.board.CountBombsAround(tc.cursor.x, tc.cursor.y)
		if got != tc.exp {
			t.Errorf("got %d but expected %d", got, tc.exp)
		}
	}
}

func TestExpandNeighbors(t *testing.T) {
	for _, tc := range []struct {
		name       string
		board      *Board
		bombs      []Point
		expandFrom Point
		expected   [][]CellState
	}{
		{
			"no bombs",
			NewBoard(3, 3),
			[]Point{},
			point(0, 0),
			[][]CellState{
				{CellStateOpened, CellStateOpened, CellStateOpened},
				{CellStateOpened, CellStateOpened, CellStateOpened},
				{CellStateOpened, CellStateOpened, CellStateOpened},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			putAllBombs(tc.board, tc.bombs)
			tc.board.expandNeighbors(tc.expandFrom.x, tc.expandFrom.y)
		})
	}
}
