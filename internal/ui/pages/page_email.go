package pages

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"go-tts/internal/ui/widgets"
)

type EmailPage struct {
	Theme      *material.Theme
	Editor     widget.Editor
	CopyBtn    widget.Clickable
	Datepicker *widgets.Datepicker
}

func NewEmailPage(th *material.Theme) *EmailPage {
	return &EmailPage{
		Theme:      th,
		Datepicker: widgets.NewDatepicker(),
	}
}

func (p *EmailPage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.Datepicker.Layout(gtx, p.Theme)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			e := material.Editor(p.Theme, &p.Editor, "Email")
			e.ReadOnly = true
			return e.Layout(gtx)
		}),
		layout.Rigid(
			material.Button(p.Theme, &p.CopyBtn, "Copy to Clipboard").Layout,
		),
	)
}
