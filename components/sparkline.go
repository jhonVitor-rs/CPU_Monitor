package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func NewSparkline(title string, lineColor, titleColor ui.Color) *widgets.Sparkline {
	sp := widgets.NewSparkline()
	sp.Title = title
	sp.LineColor = lineColor
	sp.TitleStyle.Fg = titleColor

	return sp
}

func NewSparklineGroup(title string, sp *widgets.Sparkline) *widgets.SparklineGroup {
	spg := widgets.NewSparklineGroup(sp)
	spg.Title = title
	spg.SetRect(35, 3, 80, 10)

	return spg
}
