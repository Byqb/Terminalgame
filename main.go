package main

import (
	"bytes"
	"fmt"
)

const (
	notihng = 0
	wall    = 1
	Plyer   = 69
)

type game struct {
	isRunning bool
	level *level

	drawBuf *bytes.Buffer
}

type level struct {
	width , higth int 
	data       [][]byte
}

func newLevel(width , higth int) *level {
	data := make([][]byte, higth)
	for h := 0; h < higth; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]byte, width)
		}
	}
	for h := 0; h < higth; h++ {
		for w := 0; w < width; w++ {
			if h == 0 {
				data[h][w] = wall
			}
			if w == 0 {
				data[h][w] = wall
			}
			if w == width-1 {
				data[h][w] = wall
			}
			if h == higth-1 {
				data[h][w] = wall
			}
		}
		return &level{
			width: width,
            higth: higth,
            data:  data,
		}
	}
}

func (l *level) x {} {}

func newGame(width, higth int) *game {
	lvl := newLevel(width, higth)
	return &game{}
	level: lvl,
}
func (g *game) start() {
	g.isRunning = true
	g.loop()
}

func (g *game) loop() {
	for g.isRunning {
		g.update()
		g.render()
	}
}
func (g *game) update() {}  

func (g *game) renderlevel() {

}

func (g *game) render() {}

// listen inputs
// blocking
// os.Stdin()

func main() {
	width := 80
	higth := 20 
	buf := new(bytes.Buffer)
	for h := 0; h < higth; h++ {
		for w := 0; w < width; w++ {
			if level[h][w] == notihng {
				buf.WriteString(" ")
			}
			if level[h][w] == wall {
				buf.WriteString("□")
			}
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf.String())
}
