package ui

import (
	"go-tts/internal/service"
	"go-tts/internal/ui/pages"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type UI struct {
	Theme       *material.Theme
	CurrentPage int
	TTSPage     *pages.TTSPage
}

func NewUI(ttsService *service.SyncService) *UI {
	th := material.NewTheme()
	return &UI{
		Theme:   th,
		TTSPage: pages.NewTTSPage(th),
	}
}

func (u *UI) Run(w *app.Window) error {
	var ops op.Ops

	// v0.9.0 Event Loop
	for {
		// w.Event() blocks until an event occurs
		switch e := w.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// app.NewContext creates the layout context (replaces layout.NewContext)
			gtx := app.NewContext(&ops, e)

			u.Layout(gtx)

			// Render the frame
			e.Frame(gtx.Ops)
		}
	}
}

func (u *UI) Layout(gtx layout.Context) layout.Dimensions {
	return u.TTSPage.Layout(gtx)
}
