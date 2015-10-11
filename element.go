package ansi

type Element struct {
	Rune  rune
	Color *Color
}

func NewElement() *Element {
	e := &Element{
		Rune:  ' ',
		Color: NewColor(),
	}
	return e
}

func (e *Element) String() string {
	return e.Color.String() + string(e.Rune)
}
