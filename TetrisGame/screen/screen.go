package screen

import (
	"fmt"

	"github.com/nsf/go/termbox "
)

var colors = []termbox.Attribute{
	terbox.ColorBlack,
	terbox.ColorBlue,
	terbox.ColorCyan,
	terbox.ColorWhite,
	terbox.ColorYellow,
	terbox.ColorMagenta,
	terbox.ColorLightGray,
	terbox.ColorRed,
}


type gameScreen struct{}

func (g *gameScreen)RenderAsci(board [][]int){
	fmt.Println("\n===========")
	for _, e := range board {
		for _, num :=  range e {
			if num > 0 {
				fmt.Println("X")
			} else {
				fmt.Println(" ")
			}
	}
	fmt.Println("")
	}
}
func (g *gameScreen) Render(board [][]int){
	offsety := 4
	offsetx := 4
	cellWidth := 2
	termbox.Clear(termbox.ColorGreen, termbox.ColorGreen)

	for y, e := range board {
		for x, num := range e {
		color := colors[num]
		for i := 0 ; i < cellWidth ; i ++{
		termbox.SetCell(offsetx+cellWidth*x+i,offsety+y, ' ', color, color)
		}
		}
	}

	termbox.Flush()
}

func New() *gameScreen {
	return &gameScreen{}
}
