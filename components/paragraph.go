package components

import (
	"cpu-monitor/services"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

//Criação do cabeçalho
func Header() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = "PRESS q TO QUIT"
	p.SetRect(0, 0, 80, 3)
	p.TextStyle.Fg = ui.ColorBlue
	p.TitleStyle.Fg = ui.ColorRed
	p.BorderStyle.Fg = ui.ColorBlue

	text := "Welcome to CPU Monitor!!!"
	// Calcula o deslocamento para centralizar o texto
	offsetX := (p.Inner.Dx() - len(text)) / 2
	textCenter := strings.Repeat(" ", offsetX) + text
	p.Text = textCenter

	return p
}

//Criação da caixa de texto com as informaçoes do host
func InfoHost() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = "Host Information"
	p.SetRect(0, 3, 35, 10)
	p.TextStyle.Fg = ui.ColorGreen
	p.BorderStyle.Fg = ui.ColorGreen

	//Busca das informações
	infoStat, err := services.InfoStat()
	if err != nil {
		p.Text = err.Error()
	}
	p.Text = "Operational system: " + infoStat.Hostname +
		"\nPlatform: " + infoStat.Platform +
		"\nPlatform Family: " + infoStat.PlatformFamily +
		"\nPlatform Version: " + infoStat.PlatformVersion

	return p
}
