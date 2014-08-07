package tbutil

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
)

type tbScreenPoint struct {
	tbcell termbox.Cell
	Line, Column int
}

func (p tbScreenPoint*) GetRune() rune {
	return p.tbcell.Ch
}

type tbscreen_data struct {
	Cell                                           []termbox.Cell
	visibleStarline, visibleCount, visibleMaxWidth int
}

type tbscreen struct {
	height, width int
	title         string
	data          tbscreen_data
}

func tbDrawCell(line, column int, c rune) {
	termbox.SetCell(column, line, c, termbox.ColorDefault, termbox.ColorDefault)
}

func tbDrawLine(line int, data string) {
	col := 0
	for _, x := range data {
		tbDrawCell(line, col, x)
		col++
	}
}
