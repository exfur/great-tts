package pages

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"go-tts/internal/ui/widgets"
	"gioui.org/widget"
)

type TTSPage struct {
	Theme    *material.Theme
	Table    *widgets.EditableTable
	SaveBtn  widget.Clickable
	SyncBtn  widget.Clickable
}

func NewTTSPage(th *material.Theme) *TTSPage {
	return &TTSPage{
		Theme: th,
		Table: &widgets.EditableTable{
			ListState: layout.List{Axis: layout.Vertical},
			Rows: []*widgets.TableRow{
				{
					TaskEditor:    widget.Editor{SingleLine: true, Submit: true},
					CommentEditor: widget.Editor{SingleLine: true, Submit: true},
					FromEditor:    widgets.TimeEditor{Editor: widget.Editor{SingleLine: true, Submit: true}},
					ToEditor:      widgets.TimeEditor{Editor: widget.Editor{SingleLine: true, Submit: true}},
				},
			},
		},
	}
}

func (p *TTSPage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return p.Table.Layout(gtx, p.Theme)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, material.Button(p.Theme, &p.SaveBtn, "Save").Layout),
				layout.Flexed(1, material.Button(p.Theme, &p.SyncBtn, "Sync to Jira").Layout),
			)
		}),
	)
}
