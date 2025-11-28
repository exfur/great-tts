package ui

import (
	"fmt"
	"time"
	"go-tts/internal/repository"
	"go-tts/internal/service"
	"go-tts/internal/ui/pages"

	"gioui.org/app"
	"gioui.org/io/clipboard"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type UI struct {
	Theme        *material.Theme
	CurrentPage  int
	TTSPage      *pages.TTSPage
	EmailPage    *pages.EmailPage
	NavButtons   []widget.Clickable
	ttsRepo      *repository.TTSRepository
	syncService  *service.SyncService
	emailService *service.EmailService
}

func NewUI(syncService *service.SyncService, ttsRepo *repository.TTSRepository, emailService *service.EmailService) *UI {
	th := material.NewTheme()
	return &UI{
		Theme:        th,
		TTSPage:      pages.NewTTSPage(th),
		EmailPage:    pages.NewEmailPage(th),
		NavButtons:   make([]widget.Clickable, 2),
		ttsRepo:      ttsRepo,
		syncService:  syncService,
		emailService: emailService,
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
	// Handle TTS Page button clicks
	if u.TTSPage.SaveBtn.Clicked(gtx) {
		u.ttsRepo.SaveAll()
	}
	if u.TTSPage.SyncBtn.Clicked(gtx) {
		u.syncService.Sync()
	}

	// Handle Email Page button clicks
	if u.EmailPage.Datepicker.GenerateBtn.Clicked(gtx) {
		year := u.EmailPage.Datepicker.YearEditor.Text()
		month := u.EmailPage.Datepicker.MonthEditor.Text()
		day := u.EmailPage.Datepicker.DayEditor.Text()
		dateStr := fmt.Sprintf("%s-%s-%s", year, month, day)
		// Basic validation, you might want to improve this
		if _, err := time.Parse("2006-1-2", dateStr); err == nil {
			emailBody, err := u.emailService.GenerateReportForDate(dateStr)
			if err != nil {
				// Handle error, maybe show a dialog or a message in the UI
				u.EmailPage.Editor.SetText("Error generating report: " + err.Error())
			} else {
				u.EmailPage.Editor.SetText(emailBody)
			}
		} else {
			u.EmailPage.Editor.SetText("Invalid date format. Please use YYYY-MM-DD.")
		}
	}
	if u.EmailPage.CopyBtn.Clicked(gtx) {
		gtx.Execute(clipboard.WriteCmd{Data: []byte(u.EmailPage.Editor.Text())})
	}

	// Handle navigation
	if u.NavButtons[0].Clicked(gtx) {
		u.CurrentPage = 0
	}
	if u.NavButtons[1].Clicked(gtx) {
		u.CurrentPage = 1
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, material.Button(u.Theme, &u.NavButtons[0], "TTS").Layout),
				layout.Flexed(1, material.Button(u.Theme, &u.NavButtons[1], "Email").Layout),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			switch u.CurrentPage {
			case 0:
				return u.TTSPage.Layout(gtx)
			case 1:
				return u.EmailPage.Layout(gtx)
			default:
				return layout.Dimensions{}
			}
		}),
	)
}
