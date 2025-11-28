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
	Clickable     widget.Clickable
}

type Column int

const (
	TaskColumn Column = iota
	CommentColumn
	FromColumn
	ToColumn
)

type EditableTable struct {
	ListState    layout.List
	Rows         []*TableRow
	Columns      []Column
	OnRowChanged func(rowIndex int)
	SelectedRow  int
}

func NewEditableTable(columns []Column) *EditableTable {
	return &EditableTable{
		Columns: columns,
	}
}

func (t *EditableTable) AddRow() {
	t.Rows = append(t.Rows, &TableRow{
		TaskEditor:    widget.Editor{SingleLine: true, Submit: true},
		CommentEditor: widget.Editor{SingleLine: true, Submit: true},
		FromEditor:    TimeEditor{Editor: widget.Editor{SingleLine: true, Submit: true}},
		ToEditor:      TimeEditor{Editor: widget.Editor{SingleLine: true, Submit: true}},
	})
}

func (t *EditableTable) DeleteRow(rowIndex int) {
	if rowIndex >= 0 && rowIndex < len(t.Rows) {
		t.Rows = append(t.Rows[:rowIndex], t.Rows[rowIndex+1:]...)
	}
}

func (t *EditableTable) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return t.ListState.Layout(gtx, len(t.Rows), func(gtx layout.Context, rowIndex int) layout.Dimensions {
		// Safety check: ListState might try to layout an index that was just deleted if state hasn't synced
		if rowIndex < 0 || rowIndex >= len(t.Rows) {
			return layout.Dimensions{}
		}

		row := t.Rows[rowIndex]

		// 1. Handle Selection
		if row.Clickable.Clicked(gtx) {
			t.SelectedRow = rowIndex
		}

		// 2. Event Handling Loop
		// We check for SubmitEvent (Enter key) on all editors.
		// We define a helper to reduce code duplication.
		checkForSubmit := func(editor *widget.Editor) {
			for _, e := range editor.Events() {
				if _, ok := e.(widget.SubmitEvent); ok {
					if t.OnRowChanged != nil {
						t.OnRowChanged(rowIndex)
					}
					// Move focus or handle specific logic here if needed
				}
			}
		}

		checkForSubmit(&row.TaskEditor)
		checkForSubmit(&row.CommentEditor)
		checkForSubmit(&row.FromEditor.Editor)
		checkForSubmit(&row.ToEditor.Editor)

		// 3. Layout (Drawing)
		return row.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			children := make([]layout.FlexChild, 0, len(t.Columns))

			for _, col := range t.Columns {
				// Capture loop variable
				c := col
				switch c {
				case TaskColumn:
					children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &row.TaskEditor, "Task").Layout(gtx)
					}))
				case CommentColumn:
					children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &row.CommentEditor, "Comment").Layout(gtx)
					}))
				case FromColumn:
					children = append(children, layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
						return row.FromEditor.Layout(gtx, th)
					}))
				case ToColumn:
					children = append(children, layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
						return row.ToEditor.Layout(gtx, th)
					}))
				}
			}
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx, children...)
		})
	})
}
