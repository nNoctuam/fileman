package internal

import (
	"github.com/nsf/termbox-go"
)

type Screen struct {
	header   Header
	noise    Noise
	filelist *Filelist
	footer   Footer
}

func NewScreen(filelist *Filelist) *Screen {
	return &Screen{
		header:   Header{&Area{}},
		filelist: filelist,
		footer:   Footer{&Area{}},
	}
}

func (s *Screen) Render() {
	s.filelist.Open("/home/iv")

	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.header.Area.Set(0, 0, w, 1)
	s.header.render()
	s.filelist.Area.Set(0, 1, w, h-2)
	s.filelist.render()
	s.footer.Area.Set(0, h-1, w, 1)
	s.footer.render()

	termbox.Flush()
}
