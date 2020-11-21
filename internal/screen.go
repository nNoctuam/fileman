package internal

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Area struct {
	x, y int
	w, h int
}

func NewArea(x int, y int, w int, h int) Area {
	return Area{x: x, y: y, w: w, h: h}
}

type Screen struct {
	header Header
	noise  Noise
}

func NewScreen() *Screen {
	return &Screen{
		header: Header{},
		noise:  Noise{},
	}
}

func (s *Screen) Render() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.header.renderIn(NewArea(0, 0, w, 1))
	s.noise.renderIn(NewArea(0, 1, w, h-1))

	termbox.Flush()
}

type Noise struct {
}

func (n *Noise) renderIn(area Area) {
	for y := area.y; y < area.h; y++ {
		for x := area.x; x < area.w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault,
				termbox.Attribute(rand.Int()%8)+1)
		}
	}
}

type Header struct {
}

func (h *Header) renderIn(area Area) {
	for x := area.x; x <= area.x+area.w; x++ {
		termbox.SetCell(x, area.y, '.', termbox.ColorWhite, termbox.ColorBlue)
	}
	write(time.Now().Format("2006-01-02 15:04:05"), area.x, area.y, area.w, termbox.ColorWhite, termbox.ColorBlue)
}

func write(text string, x, y, w int, fg, bg termbox.Attribute) {
	for pos := x; pos-x < len(text); pos++ {
		termbox.SetCell(pos, y, rune(text[pos]), fg, bg)
	}
}
