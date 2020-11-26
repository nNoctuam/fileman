package internal

import (
	"github.com/nsf/termbox-go"
)

type Screen struct {
	header Header
	noise  Noise
	footer Footer
}

func NewScreen() *Screen {
	return &Screen{
		header: Header{&Area{}},
		noise:  Noise{&Area{}},
		footer: Footer{&Area{}},
	}
}

func (s *Screen) Render() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.header.Area.Set(0, 0, w, 1)
	s.header.render()
	s.noise.Area.Set(0, 1, w, h-1)
	s.noise.render()
	s.footer.Area.Set(0, h-1, w, 1)
	s.footer.render()

	termbox.Flush()
}
