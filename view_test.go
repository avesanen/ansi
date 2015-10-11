package ansi

import (
	"fmt"
	"testing"
)

func TestWindowGetElement(t *testing.T) {
	w := 80
	h := 25
	w1 := NewWindow()
	w1.W = w
	w1.H = h
	w1.Decorated = true
	canvas := NewCanvas(w, h)
	canvas.GetElement(10, 10).Rune = 'X'
	canvas.GetElement(10, 10).Color.Fg = 4
	w1.Draw = canvas

	s := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			s += w1.GetElement(x, y).String()
		}
		s += "\n"
	}
	fmt.Print(s)
}
