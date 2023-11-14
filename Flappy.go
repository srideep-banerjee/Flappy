package main

import (
	"fmt"
	term "github.com/nsf/termbox-go"
)

type size struct {
	height int;
	width int;
}


func main() {
	err:=term.Init()
	if (err != nil) {
		fmt.Println("Terminal does not support required library: termbox-go")
		panic(err)
	}
	defer term.Close()
	
	//width,heigth :=term.Size()
	
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
	term.Flush()
}

func drawCanvas() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 100; j++ {
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
	
}