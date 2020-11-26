package internal

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Footer struct {
	Area *Area
}

func (f *Footer) render() {
	for x := f.Area.X; x <= f.Area.X+f.Area.W; x++ {
		termbox.SetCell(x, f.Area.Y, '.', termbox.ColorWhite, termbox.ColorBlue)
	}
	f.Area.write(time.Now().Format("2006-01-02 15:04:05"), f.Area.X, f.Area.Y, f.Area.W, termbox.ColorWhite, termbox.ColorBlue)
}
