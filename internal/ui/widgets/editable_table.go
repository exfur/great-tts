package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TableRow struct {
	TaskEditor    widget.Editor
	CommentEditor widget.Editor
	FromEditor    TimeEditor
	ToEditor      TimeEditor
	// ... holds state for one row
}

type EditableTable struct {
	ListState layout.List
	Rows      []*TableRow // Persistent state
}

func (t *EditableTable) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return t.ListState.Layout(gtx, len(t.Rows), func(gtx layout.Context, rowIndex int) layout.Dimensions {
		row := t.Rows[rowIndex]
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Editor(th, &row.TaskEditor, "Task").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Editor(th, &row.CommentEditor, "Comment").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return row.FromEditor.Layout(gtx, th)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return row.ToEditor.Layout(gtx, th)
			}),
		)
	})
}
