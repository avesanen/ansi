package ansi

type Draw interface {
	GetElement(x, y int) *Element
}

type Window struct {
	X, Y, W, H int
	Title      string
	Decorated  bool
	SubWindows []*Window
	Draw       Draw
}

func NewWindow() *Window {
	w := &Window{}
	return w
}

func (w *Window) GetElement(x, y int) *Element {
	for _, s := range w.SubWindows {
		if x >= s.X && y >= s.Y && x < s.X+s.W && y < s.Y+s.H {
			return s.GetElement(x, y)
		}
	}
	if w.Draw != nil {
		return w.Draw.GetElement(x, y)
	} else {
		return &Element{Rune: '.'}
	}
}

/*
TODO: Something like this:

func (v *View) decoration(x, y int) *Element {
	if y == v.y && x > v.x && x < v.x+len(v.name)+3 {
		if x == v.x+1 {
			return &Element{Rune: '['}
		} else if x == v.x+len(v.name)+2 {
			return &Element{Rune: ']'}
		}
		return &Element{Rune: rune(v.name[x-v.x-2])}
	}
	if x == v.x && y == v.y {
		return &Element{Rune: '+'}
	} else if x == v.x+v.w-1 && y == v.y+v.h-1 {
		return &Element{Rune: '+'}
	} else if x == v.x && y == v.y+v.h-1 {
		return &Element{Rune: '+'}
	} else if x == v.x+v.w-1 && y == v.y {
		return &Element{Rune: '+'}
	} else if x == v.x {
		return &Element{Rune: '|'}
	} else if y == v.y {
		return &Element{Rune: '-'}
	} else if x == v.x+v.w-1 {
		return &Element{Rune: '|'}
	} else if y == v.y+v.h-1 {
		return &Element{Rune: '-'}
	}
	return nil
}
*/
