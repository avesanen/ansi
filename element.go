package ansi

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// Element is one "pixel" of the view
type Element struct {
	Rune    rune
	Bold    bool
	FgColor int
	BgColor int
}

// fgBytes returns the foreground color as ansi escape code
func (e *Element) fgBytes() []byte {
	if e.Bold {
		return []byte(fmt.Sprintf("\033[0;3" + strconv.Itoa(e.FgColor) + "m"))
	} else {
		return []byte(fmt.Sprintf("\033[1;3" + strconv.Itoa(e.FgColor) + "m"))
	}
}

// bgBytes returns the background color as ansi escape code
func (e *Element) bgBytes() []byte {
	return []byte(fmt.Sprintf("\033[4" + strconv.Itoa(e.BgColor) + "m"))
}

// bytes returns the element as ansi bytes with color escape codes and the utf-8 character
func (e *Element) bytes() []byte {
	// TODO: Slow and gc heavy baybe, so optimize later maybe.
	rb := make([]byte, utf8.RuneLen(e.Rune))
	utf8.EncodeRune(rb, e.Rune)
	b := e.fgBytes()
	b = append(b, e.bgBytes()...)
	b = append(b, rb...)
	return b
}
