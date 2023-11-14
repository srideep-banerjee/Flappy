package main

import (
	"fmt"
	term "github.com/nsf/termbox-go"
)

type Size struct {
	height int;
	width int;
}

var screenSize Size = Size{height: 25, width: 100}

func main() {
	err:=term.Init()
	if (err != nil) {
		fmt.Println("Terminal does not support required library: termbox-go")
		panic(err)
	}
	defer term.Close()
	
	screenSize.width,screenSize.height = term.Size()
	
	render()
	
	var event = term.PollEvent()
	var location = 0
	for event.Key != term.KeyEnter {
		term.SetChar(location, 26, event.Ch)
		term.Flush()
		location++
		event = term.PollEvent()
	}
}

func render() {
	drawCanvas()
	drawBird(10, 10)
	drawPipe(20, 10, true)
	term.Flush()
}

func drawCanvas() {
	for i := 0; i < screenSize.height; i++ {
		for j := 0; j < screenSize.width; j++ {
			term.SetCell(j, i, ' ', term.ColorBlack, term.ColorCyan)
		}
	}
}

func drawBird(x, y int) {
	term.SetCell(x, y, ' ', term.ColorDefault, term.ColorYellow)
	term.SetCell(x + 1, y, '\u2022', term.ColorBlack, term.ColorYellow)
	term.SetCell(x + 2, y, '-', term.ColorRed, term.ColorCyan)
}

func drawPipe(x, y int, down bool) {
	if (down) {
		for i := y; i <= screenSize.height; i++ {
			for j := x; j < x + 4; j++ {
				term.SetCell(j, i, ' ', term.ColorDefault, term.ColorGreen)
			}
		}
	}
}