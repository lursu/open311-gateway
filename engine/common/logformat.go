package common

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	// "github.com/davecgh/go-spew/spew"
)

// ==============================================================================================================================
//                                      CONSOLE
// ==============================================================================================================================

// NewLogString creates a new LogString, and initializes color printing.
func NewLogString() *LogString {
	ls := new(LogString)
	ls.color = make(map[string]func(...interface{}) string)
	ls.color["red"] = color.New(color.FgRed).SprintFunc()
	ls.color["green"] = color.New(color.FgGreen).SprintFunc()
	ls.color["blue"] = color.New(color.FgBlue).SprintFunc()
	ls.color["yellow"] = color.New(color.FgYellow).SprintFunc()
	return ls
}

// LogString is used to "box" object representations.
type LogString struct {
	raw   string
	fmt   string
	color map[string]func(...interface{}) string
}

// AddF adds a formated line of text, like Printf().
func (l *LogString) AddF(format string, args ...interface{}) {
	l.raw = l.raw + fmt.Sprintf(format, args...)
}

// AddS adds a single line of text, with no terminating line return.
func (l *LogString) AddS(s string) {
	l.raw = l.raw + s
}

// AddSR adds a single line of text (with line return), like Println().
func (l *LogString) AddSR(s string) {
	l.raw = l.raw + s + "\n"
}

// Color applies the specified color to the string.
func (l *LogString) Color(s, color string) string {
	f, ok := l.color[color]
	if !ok {
		return s
	}
	return f(s)
}

// ColorBool returns a colored string depending on whether the specified value is true or false.
func (l *LogString) ColorBool(v bool, strue, sfalse, ctrue, cfalse string) string {
	if v {
		return l.Color(strue, ctrue)
	}
	return l.Color(sfalse, cfalse)
}

// Box draws a box around the LogString with the specified line width, with a leading line return.
func (l *LogString) Box(w int) string {
	return l.box(w, true)
}

// BoxC draws a box around the LogString with the specified line width, without a leading line return.
func (l *LogString) BoxC(w int) string {
	return l.box(w, false)
}

// box draws a box around the LogString with the specified line width, and leading line return.
func (l *LogString) box(w int, lr bool) string {
	var out string
	if lr {
		out = "\n"
	}
	ss := strings.Split(l.raw, "\n")
	ls := len(ss)
	for i, ln := range ss {
		if i == 0 {
			x := ((w - len(ln)) / 2) - 1
			out += fmt.Sprintf("\u2554%s %s %s\n", strings.Repeat("\u2550", x), ln, strings.Repeat("\u2550", x))
		} else if i == (ls-1) && len(ln) == 0 {
			continue
		} else {
			out += fmt.Sprintf("\u2551%s\n", strings.Replace(ln, "\n", "\n\u2551", -1))
		}
	}
	out += fmt.Sprintf("\u255A%s\n", strings.Repeat("\u2550", w))
	l.fmt = out
	return l.fmt
}
