package ansi

type Cursor struct {
	X, Y  int
	Color *Color
}

func NewCursor() *Cursor {
	c := &Cursor{
		Color: NewColor(),
	}
	return c
}
