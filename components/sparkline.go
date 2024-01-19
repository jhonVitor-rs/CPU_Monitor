package components

import (
	"github.com/gizak/termui/v3/widgets"
)

func NewSparkline() *widgets.Sparkline {
	sp := widgets.NewSparkline()

	return sp
}

func NewSparklineGroup(title string, sp *widgets.Sparkline) *widgets.SparklineGroup {
	spg := widgets.NewSparklineGroup(sp)
	spg.Title = title
	spg.SetRect(35, 3, 80, 12)

	return spg
}
