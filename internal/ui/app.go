package ui

import (
	"fmt"
	"go-tts/internal/model"
	"go-tts/internal/repository"
	"go-tts/internal/service"
	"go-tts/internal/ui/pages"
	"io"
	"strings"
	"time"

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
	RegistryPage *pages.RegistryPage
	NavButtons   []widget.Clickable
	ttsRepo      repository.TTSRepository
	registryRepo repository.RegistryRepository
	syncService  *service.SyncService
	emailService *service.EmailService
}

func NewUI(syncService *service.SyncService, ttsRepo repository.TTSRepository, registryRepo repository.RegistryRepository, emailService *service.EmailService) *UI {
	th := material.NewTheme()
	ttsPage := pages.NewTTSPage(th)
	emailPage := pages.NewEmailPage(th)
	registryPage := pages.NewRegistryPage(th)

	ui := &UI{
		Theme:        th,
		TTSPage:      ttsPage,
		EmailPage:    emailPage,
		RegistryPage: registryPage,
		NavButtons:   make([]widget.Clickable, 3),
		ttsRepo:      ttsRepo,
		registryRepo: registryRepo,
		syncService:  syncService,
		emailService: emailService,
	}

	ttsPage.Table.OnRowChanged = func(rowIndex int) {
		row := ttsPage.Table.Rows[rowIndex]
		entry := model.TTSLogEntry{
			Task:    row.TaskEditor.Text(),
			Comment: row.CommentEditor.Text(),
			From:    row.FromEditor.Editor.Text(),
			To:      row.ToEditor.Editor.Text(),
		}
		ttsRepo.Save(entry)
	}

	registryPage.Table.OnRowChanged = func(rowIndex int) {
		row := registryPage.Table.Rows[rowIndex]
		entry := model.RegistryEntry{
			Task:  row.TaskEditor.Text(),
			Issue: row.CommentEditor.Text(),
		}
		registryRepo.Save(entry)
	}

	return ui
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
	if u.TTSPage.AddBtn.Clicked(gtx) {
		u.TTSPage.Table.AddRow()
	}
	if u.TTSPage.DeleteBtn.Clicked(gtx) {
		u.TTSPage.Table.DeleteRow(u.TTSPage.Table.SelectedRow)
	}
	if u.TTSPage.SyncBtn.Clicked(gtx) {
		u.syncService.SyncApprovedWork()
	}

	// Handle Registry Page button clicks
	if u.RegistryPage.AddBtn.Clicked(gtx) {
		u.RegistryPage.Table.AddRow()
	}
	if u.RegistryPage.DeleteBtn.Clicked(gtx) {
		u.RegistryPage.Table.DeleteRow(u.RegistryPage.Table.SelectedRow)
	}

	// Handle Email Page button clicks
	if u.EmailPage.Datepicker.GenerateBtn.Clicked(gtx) {
		year := u.EmailPage.Datepicker.YearEditor.Text()
		month := u.EmailPage.Datepicker.MonthEditor.Text()
		day := u.EmailPage.Datepicker.DayEditor.Text()
		dateStr := fmt.Sprintf("%s-%s-%s", year, month, day)
		// Basic validation, you might want to improve this
		if date, err := time.Parse("2006-1-2", dateStr); err == nil {
			emailBody, err := u.emailService.GenerateReport(date)
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
		reader := strings.NewReader(u.EmailPage.Editor.Text())
		gtx.Execute(clipboard.WriteCmd{Data: io.NopCloser(reader)})
	}

	// Handle navigation
	if u.NavButtons[0].Clicked(gtx) {
		u.CurrentPage = 0
	}
	if u.NavButtons[1].Clicked(gtx) {
		u.CurrentPage = 1
	}
	if u.NavButtons[2].Clicked(gtx) {
		u.CurrentPage = 2
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, material.Button(u.Theme, &u.NavButtons[0], "TTS").Layout),
				layout.Flexed(1, material.Button(u.Theme, &u.NavButtons[1], "Email").Layout),
				layout.Flexed(1, material.Button(u.Theme, &u.NavButtons[2], "Registry").Layout),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			switch u.CurrentPage {
			case 0:
				return u.TTSPage.Layout(gtx)
			case 1:
				return u.EmailPage.Layout(gtx)
			case 2:
				return u.RegistryPage.Layout(gtx)
			default:
				return layout.Dimensions{}
			}
		}),
	)
}
