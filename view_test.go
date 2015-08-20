package ansi

import (
	"testing"
)

func TestViewMove(t *testing.T) {
	v := &View{x: 0, y: 0, w: 10, h: 10}
	v.move(1, 2)
	if v.x != 1 || v.y != 2 {
		t.Logf("{x:%d,y:%d} expected {x:1,y:2}", v.x, v.y)
		t.Fail()
	}
}

func TestViewResize(t *testing.T) {
	v := &View{x: 0, y: 0, w: 10, h: 10}
	v.resize(11, 12)
	if v.w != 11 || v.h != 12 {
		t.Logf("{w:%d,h:%d} expected {w:11,y:12}", v.x, v.y)
		t.Fail()
	}
}

func BuildViewWithSubs() *View {
	v := &View{x: 0, y: 0, w: 10, h: 10}
	v.name = "parent"
	sv1 := v.newSub(0, 0, 5, 10)
	sv1.name = "sv1"
	sv2 := v.newSub(5, 0, 5, 10)
	sv2.name = "sv2"
	return v
}

func TestViewNewSub(t *testing.T) {
	v := BuildViewWithSubs()
	if len(v.subViews) != 2 {
		t.Logf("len(subViews)==%d expected len(subViews)==2", len(v.subViews))
		t.Fail()
	}
}

func TestViewGetSubAt(t *testing.T) {
	v := BuildViewWithSubs()

	getSv1 := v.getSubAt(4, 5)
	getSv2 := v.getSubAt(5, 5)

	if getSv1.name != "sv1" {
		t.Logf("subAt(4,5)==\"%s\" expected subAt(4,5)==\"%s\"", getSv1.name, "sv1")
		t.Fail()
	}
	if getSv2.name != "sv2" {
		t.Logf("subAt(5,5)==\"%s\" expected subAt(5,5)==\"%s\"", getSv2.name, "sv2")
		t.Fail()
	}
}

func TestViewRmSub(t *testing.T) {
	v := BuildViewWithSubs()
	getSv2a := v.getSubAt(5, 5)
	v.rmSub(getSv2a)
	getSv2b := v.getSubAt(5, 5)
	if getSv2a == getSv2b {
		t.Logf("subAt(5,5)==\"%s\" expected subAt(5,5)==\"%s\"", getSv2b.name, "parent")
	}
}

type TestDraw struct {
	Rune    rune
	FgColor int
	BgColor int
}

func (td *TestDraw) GetElementAt(x, y int) *Element {
	return &Element{Rune: td.Rune, FgColor: td.FgColor, BgColor: td.BgColor}
}

func (td *TestDraw) Dirty() bool { return true }

func TestViewGetElement(t *testing.T) {
	v := BuildViewWithSubs()
	getSv2a := v.getSubAt(4, 5)
	getSv2b := v.getSubAt(5, 5)
	getSv2a.draw = &TestDraw{Rune: 'A', FgColor: 4, BgColor: 2}
	getSv2b.draw = &TestDraw{Rune: 'B', FgColor: 5, BgColor: 3}

	e1 := v.GetElement(4, 5)
	if e1.Rune != 'A' {
		t.Logf("GetElement(4,5)==\"%b\" expected GetElement(4,5)==\"%b\"", e1.Rune, 'A')
		t.Fail()
	}

	e2 := v.GetElement(5, 5)
	if e2.Rune != 'B' {
		t.Logf("GetElement(5,5)==\"%b\" expected GetElement(5,5)==\"%b\"", e2.Rune, 'B')
		t.Fail()
	}
}
