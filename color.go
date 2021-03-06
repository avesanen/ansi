package ansi

import "strconv"

type Color struct {
	Fg   int
	Bg   int
	Bold bool
}

func NewColor() *Color {
	c := &Color{
		Fg:   7,
		Bg:   0,
		Bold: false,
	}
	return c
}

func (c *Color) String() string {
	s := "\033["
	if c.Bold {
		s += "1;"
	} else {
		s += "0;"
	}
	s += "3" + strconv.Itoa(c.Fg)
	s += ";4" + strconv.Itoa(c.Bg)
	s += "m"
	return s
}
