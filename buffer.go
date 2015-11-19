package ansi

import (
	"strconv"
	"strings"
)

type Buffer struct {
	Window    *Window
	Buffer    *Canvas
	Outbuffer string
	Cursor    *Cursor
}

func NewBuffer(w *Window) *Buffer {
	b := &Buffer{
		Window: w,
		Buffer: NewCanvas(w.W, w.H),
		Cursor: NewCursor(),
	}
	return b
}

func (b *Buffer) Refresh() string {
	b.Outbuffer = ""
	moved := true
	for y := 0; y < b.Window.H; y++ {
		for x := 0; x < b.Window.W; x++ {
			changed := false
			winEle := b.Window.GetElement(x, y)
			bufEle := b.Buffer.GetElement(x, y)
			if winEle.Color.Fg != bufEle.Color.Fg ||
				winEle.Color.Bg != bufEle.Color.Bg ||
				winEle.Color.Bold != bufEle.Color.Bold {
				b.SetColor(winEle.Color)
				changed = true
				bufEle.Color.Fg = winEle.Color.Fg
				bufEle.Color.Bg = winEle.Color.Bg
				bufEle.Color.Bold = winEle.Color.Bold
			}

			if winEle.Rune != bufEle.Rune || changed {
				if moved {
					b.SetPosition(x, y)
					moved = false
				}
				b.Outbuffer += string(winEle.Rune)
				bufEle.Rune = winEle.Rune
				b.Cursor.X += 1
				if b.Cursor.Y > b.Window.W {
					b.Cursor.X = 0
					b.Cursor.Y += 1
					moved = true
				}
			} else {
				moved = true
			}

		}
	}
	return b.Outbuffer
}

func (b *Buffer) SetColor(c *Color) {
	var params []string

	if b.Cursor.Color.Fg != c.Fg {
		params = append(params, "3"+strconv.Itoa(c.Fg))
		b.Cursor.Color.Fg = c.Fg
	}
	if b.Cursor.Color.Bg != c.Bg {
		params = append(params, "4"+strconv.Itoa(c.Bg))
		b.Cursor.Color.Bg = c.Bg
	}
	if b.Cursor.Color.Bold != c.Bold {
		if c.Bold {
			params = append(params, "1")
		} else {
			params = append(params, "0")
		}
		b.Cursor.Color.Bold = c.Bold
	}

	if len(params) != 0 {
		b.Outbuffer += "\033[" + strings.Join(params, ";") + "m"
	}
}

func (b *Buffer) SetPosition(x, y int) {
	b.Outbuffer += "\033[" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + "H"
	b.Cursor.X = x
	b.Cursor.Y = y
}
