package widgets

import (
	"fmt"
	"time"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Datepicker struct {
	YearEditor  widget.Editor
	MonthEditor widget.Editor
	DayEditor   widget.Editor
	GenerateBtn widget.Clickable
}

func NewDatepicker() *Datepicker {
	now := time.Now()
	dp := &Datepicker{}
	dp.YearEditor.SetText(fmt.Sprintf("%d", now.Year()))
	dp.MonthEditor.SetText(fmt.Sprintf("%d", now.Month()))
	dp.DayEditor.SetText(fmt.Sprintf("%d", now.Day()))
	return dp
}

func (dp *Datepicker) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(material.Editor(th, &dp.YearEditor, "YYYY").Layout),
		layout.Rigid(material.Editor(th, &dp.MonthEditor, "MM").Layout),
		layout.Rigid(material.Editor(th, &dp.DayEditor, "DD").Layout),
		layout.Rigid(material.Button(th, &dp.GenerateBtn, "Generate").Layout),
	)
}
