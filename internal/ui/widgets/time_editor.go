package widgets

import (
	"regexp"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var timeRegex = regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`)

type TimeEditor struct {
	widget.Editor
	OnSubmit func()
	lastText string
}

func (t *TimeEditor) Frame() {
	if t.Text() != t.lastText {
		if timeRegex.MatchString(t.Text()) {
			if t.OnSubmit != nil {
				t.OnSubmit()
			}
		}
	}
	t.lastText = t.Text()
}

func (t *TimeEditor) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	t.Frame()
	return material.Editor(th, &t.Editor, "HH:MM").Layout(gtx)
}
