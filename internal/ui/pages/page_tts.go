package pages

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"go-tts/internal/ui/widgets"
	"gioui.org/widget"
)

type TTSPage struct {
	Theme *material.Theme
	Table *widgets.EditableTable
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
	return p.Table.Layout(gtx, p.Theme)
}
