package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func NewParagraph(title string, x1, y1, x2, y2 int, textColor, borderColor ui.Color) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = title
	p.SetRect(x1, y1, x2, y2)
	p.TextStyle.Fg = textColor
	p.BorderStyle.Fg = borderColor

	return p
}
