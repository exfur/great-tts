package pages

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"go-tts/internal/ui/widgets"
)

type RegistryPage struct {
	Theme *material.Theme
	Table *widgets.EditableTable
	AddBtn   widget.Clickable
	DeleteBtn widget.Clickable
}

func NewRegistryPage(th *material.Theme) *RegistryPage {
	return &RegistryPage{
		Theme: th,
		Table: widgets.NewEditableTable([]widgets.Column{
			widgets.TaskColumn,
			widgets.CommentColumn,
		}),
	}
}

func (p *RegistryPage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return p.Table.Layout(gtx, p.Theme)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, material.Button(p.Theme, &p.AddBtn, "Add").Layout),
				layout.Flexed(1, material.Button(p.Theme, &p.DeleteBtn, "Delete").Layout),
			)
		}),
	)
}
