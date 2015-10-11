package ansi

import "fmt"

type Canvas struct {
	W, H     int
	Elements []*Element
}

func NewCanvas(w, h int) *Canvas {
	c := &Canvas{
		W: w,
		H: h,
	}
	c.Resize(w, h)
	return c
}

func (c *Canvas) Resize(w, h int) {
	if w*h <= 0 {
		return
	}
	c.Elements = make([]*Element, w*h)
	for i, _ := range c.Elements {
		c.Elements[i] = NewElement()
	}
}

func (c *Canvas) GetElement(x, y int) *Element {
	if x < 0 || y < 0 || x > c.W-1 || y > c.H-1 {
		fmt.Println("Out of bounds!", x, y, c.W, c.H)
		return nil
	}
	return c.Elements[x+y*c.W]
}
