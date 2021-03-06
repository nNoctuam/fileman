package main

import (
	"filemanager/internal"
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	filelist := internal.NewFilelist()
	err := filelist.Open("/")
	if err != nil {
		fmt.Println(err)
	}
	screen := internal.NewScreen(filelist)

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			screen.Render()
			//time.Sleep(16667 * time.Microsecond)
			time.Sleep(1 * time.Second)
		}
	}
}
