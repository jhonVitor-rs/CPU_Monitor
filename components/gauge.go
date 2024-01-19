package components

import (
	"github.com/gizak/termui/v3/widgets"
)

func MemoryMonitor() *widgets.Gauge {
	g := widgets.NewGauge()
	g.Title = "Memory Used"
	g.SetRect(0, 20, 80, 25)

	return g
}
