package pages

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type TTSPage struct {
	Theme *material.Theme
}

func NewTTSPage(th *material.Theme) *TTSPage {
	return &TTSPage{Theme: th}
}

func (p *TTSPage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.H3(p.Theme, "TTS Dashboard").Layout(gtx)
	})
}
