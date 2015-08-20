package ansi

// Draw is the interface for drawing a view
type Draw interface {
	GetElementAt(x, y int) *Element
	Dirty() bool
}

// View is a defined area with subviews.
// View can also have a draw interface
type View struct {
	name          string
	x, y, h, w, z int
	dirty         bool
	subViews      []*View
	draw          Draw
}

// move moves the view at (x,y) position
func (v *View) move(x, y int) {
	v.x = x
	v.y = y
}

// resize changes dimensions to (w,h)
func (v *View) resize(w, h int) {
	v.w = w
	v.h = h
	v.dirty = true
}

// sortSubs first sorts own subs by z, then asks all subViews to sort
func (v *View) sortSubs() {
	for i := 0; i < len(v.subViews); i++ {
		for j := 0; j < len(v.subViews); j++ {
			if v.subViews[i].z > v.subViews[j].z {
				v.subViews[i], v.subViews[j] = v.subViews[j], v.subViews[i]
			}
		}
	}
	for _, subView := range v.subViews {
		subView.sortSubs()
	}
}

// NewView returns a new view with (x,y) cordinates and (w,h) dimensions
func (v *View) newSub(x, y, w, h int) *View {
	newView := &View{x: x, y: y, h: h, w: w, z: 0}
	v.addSub(newView)
	return newView
}

// addSub adds a new subView
func (v *View) addSub(newView *View) {
	v.subViews = append(v.subViews, newView)
	v.sortSubs()
}

// rmSub removes a subView
func (v *View) rmSub(rmView *View) {
	for i, k := range v.subViews {
		if k == rmView {
			v.subViews = append(v.subViews[:i], v.subViews[i+1:]...)
		}
	}
}

// getSubAt checks there is a subview at the (x,y) position then
// returns the getSubAt of that view. If there is no subView at that
// position, that subview is then returned.
func (v *View) getSubAt(x, y int) *View {
	for _, k := range v.subViews {
		if x >= k.x && y >= k.y && x < k.x+k.w && y < k.y+k.h {
			return k.getSubAt(x-k.x, y-k.y)
		}
	}
	return v
}

// getElement returns the element from this views draw interface,
// if interface is nil, returns nil
func (v *View) getElement(x, y int) *Element {
	if v.draw == nil {
		return nil
	}
	return v.draw.GetElementAt(x, y)
}

// GetElement first gets the subview at the (x,y) position and
// then returns the element drawn by that view
func (v *View) GetElement(x, y int) *Element {
	drawnView := v.getSubAt(x, y)
	return drawnView.getElement(x, y)
}
