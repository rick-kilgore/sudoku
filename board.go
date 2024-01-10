package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type Cell struct {
	value *byte
	fixed bool
	hints []bool
}

func NewEmptyCell(boardWidth int) *Cell {
	return &Cell{
		hints: make([]bool, boardWidth),
	}
}

func NewFixedCell(boardWidth int, value byte) *Cell {
	return &Cell{
		value: &value,
		fixed: true,
		hints: make([]bool, boardWidth),
	}
}

func (c *Cell) Display() []string {
	if c.value == nil {
		return c.displayHints()
	}
	return c.displayValue()
}

func (c *Cell) displayHints() []string {
	width := int(math.Sqrt(float64(len(c.hints))))
	var rows []string
	for i := 0; i < width; i++ {
		sRow := ""
		for j := 0; j < width; j++ {
			idx := i*width + j
			if c.hints[idx] {
				sRow += fmt.Sprintf("%d", idx+1)
			} else {
				sRow += " "
			}
		}
		rows = append(rows, sRow)
	}
	return rows
}

func (c *Cell) displayValue() []string {
	width := int(math.Sqrt(float64(len(c.hints))))
	var rows []string
	for i := 0; i < width; i++ {
		sRow := ""
		for j := 0; j < width; j++ {
			if i == width/2 && j == width/2 {
				sRow += fmt.Sprintf("[1;32m%d[m", *c.value)
			} else {
				sRow += " "
			}
		}
		rows = append(rows, sRow)
	}
	return rows
}

type Board struct {
	width int
	cells []*Cell
}

func NewBoardFromFile(width int, rdr io.Reader) *Board {
	board := &Board{
		width: width,
	}
	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		sRow := scanner.Text()
		if width != len(sRow) {
			panic(fmt.Sprintf("incorrect row length - expected %d: %s", width, sRow))
		}
		for i := 0; i < width; i++ {
			if sRow[i] == ' ' {
				board.cells = append(board.cells, NewEmptyCell(width))
			} else {
				board.cells = append(board.cells, NewFixedCell(width, sRow[i]-'0'))
			}
		}
	}
	return board
}

func (b *Board) index(i, j int) int {
	return i*b.width + j
}

func (b *Board) at(i, j int) *Cell {
	return b.cells[b.index(i, j)]
}

func (b *Board) value(i, j int) *byte {
	return b.cells[b.index(i, j)].value
}

func (b *Board) hints(i, j int) []bool {
	return b.cells[b.index(i, j)].hints
}

func (b *Board) isPossible(i, j int, value byte) bool {
	return b.hints(i, j)[value-1]
}

func (b *Board) Display() []string {
	var rows []string
	sqrt := int(math.Sqrt(float64(b.width)))

	for i := 0; i < b.width; i++ {
		for n := 0; n < sqrt; n++ {
			rows = append(rows, "| ")
		}
		for j := 0; j < b.width; j++ {
			cellRows := b.at(i, j).Display()
			for n := 0; n < sqrt; n++ {
				rows[sqrt*i+n] += cellRows[n] + " | "
			}
		}
	}
	return rows
}
