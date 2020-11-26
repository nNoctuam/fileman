package internal

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Area struct {
	X, Y int
	W, H int
}

func (a *Area) Set(x int, y int, w int, h int) *Area {
	a.X = x
	a.Y = y
	a.W = w
	a.H = h
	return a
}

func (a *Area) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if x < a.X || y < a.Y || x > a.X+a.W || y > a.Y+a.H {
		fmt.Printf("cell out of bounds: %+v %d %d", a, x, y)
		return
	}
	termbox.SetCell(x, y, ch, fg, bg)
}

func (a *Area) write(text string, x, y, w int, fg, bg termbox.Attribute) {
	for pos := x; pos-x < len(text) && pos-x < w; pos++ {
		a.SetCell(pos, y, rune(text[pos]), fg, bg)
	}
}
