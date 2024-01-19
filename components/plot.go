package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CpuUsage() *widgets.Plot {
	p := widgets.NewPlot()
	p.Title = "CPU Usage"
	p.Data = make([][]float64, 1)
	p.SetRect(35, 12, 80, 20)
	p.AxesColor = ui.ColorWhite
	p.LineColors[0] = ui.ColorYellow

	return p
}
