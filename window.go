package ansi

type Draw interface {
	GetElement(x, y int) *Element
}

type Window struct {
	X, Y, W, H int
	Title      string
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
			return s.GetElement(x-s.X, y-s.Y)
		}
	}

	if w.Draw != nil {
		return w.Draw.GetElement(x, y)
	}
	return NewElement()
}
