package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nsf/termbox-go"
)

type Filelist struct {
	Path     string
	Contains []string
	Area     *Area
}

func NewFilelist() *Filelist {
	return &Filelist{
		Area: &Area{},
	}
}

func (l *Filelist) Open(path string) error {
	stat, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return fmt.Errorf("'%s' is not a directory", path)
	}

	l.Path = path
	files, err := ioutil.ReadDir(l.Path)
	if err != nil {
		return fmt.Errorf("couldn't read '%s'", path)
	}
	l.Contains = make([]string, len(files))
	for i, file := range files {
		l.Contains[i] = file.Name()
	}

	//fmt.Printf("%+v", stat)
	//time.Sleep(time.Second * 5)

	return nil
}

func (l *Filelist) render() {
	l.Area.writeLines(l.Contains, l.Area.X, l.Area.Y, l.Area.W, l.Area.H, termbox.ColorWhite, termbox.ColorDefault)
}
