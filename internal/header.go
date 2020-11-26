package internal

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Header struct {
	Area *Area
}

func (h *Header) render() {
	for x := h.Area.X; x <= h.Area.X+h.Area.W; x++ {
		termbox.SetCell(x, h.Area.Y, '.', termbox.ColorWhite, termbox.ColorBlue)
	}
	h.Area.writeLine(time.Now().Format("FILE MANAGER"), h.Area.X, h.Area.Y, h.Area.W, termbox.ColorWhite, termbox.ColorBlue)
}
