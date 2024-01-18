package main

import (
	"cpu-monitor/components"
	"cpu-monitor/services"
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	// Inicialização do Termui
	if err := ui.Init(); err != nil {
		fmt.Printf("Erro ao inicializar o termui: %v\n", err)
		return
	}
	defer ui.Close()

	header := components.Header()

	infoHost := components.InfoHost()

	process := components.Process()

	temp := components.NewSparkline("Temperature", ui.ColorCyan, ui.ColorCyan)
	tempStat := components.NewSparklineGroup("Temperature Status", temp)

	//Processos em execução
	processData := make([]string, 0, 50)

	//Criação do monitor de temperatura
	tempData := make([]float64, 0, 50)

	draw := func(count int) {
		//Preenchimento da lista de processos
		prcessInfo, _ := services.InfoProcess()
		processData = nil
		for i, proc := range prcessInfo {
			name, _ := proc.Name()
			processData = append(processData, fmt.Sprintf("[%d] - %s", i+1, name))
		}
		process.Rows = processData[count%len(processData):]

		//Preenchimento da temperatura
		temps, _ := services.TemperatureStat()
		for _, temp := range temps {
			tempData = append(tempData, temp.Temperature)
		}
		if len(tempData) > 50 {
			tempData = tempData[len(tempData)-50:]
		}
		temp.Data = tempData

		ui.Clear()
		ui.Render(header, infoHost, process, tempStat)
	}

	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}

		case <-ticker:
			draw(tickerCount)
			tickerCount++
		}
	}
}
