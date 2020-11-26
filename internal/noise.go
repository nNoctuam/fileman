package internal

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

type Noise struct {
	Area *Area
}

func (n *Noise) render() {
	for y := n.Area.Y; y < n.Area.H; y++ {
		for x := n.Area.X; x < n.Area.W; x++ {
			n.Area.SetCell(x, y, ' ', termbox.ColorDefault,
				termbox.Attribute(rand.Int()%8)+1)
		}
	}
}
