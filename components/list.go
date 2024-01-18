package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Process() *widgets.List {
	l := widgets.NewList()
	l.Title = "Process"
	l.SetRect(0, 10, 35, 20)
	l.TextStyle.Fg = ui.ColorYellow

	return l
}
