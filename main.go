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

	temp := components.NewSparkline()
	tempStat := components.NewSparklineGroup("Temperature", temp)

	cpuUsage := components.CpuUsage()

	memoryUsage := components.MemoryMonitor()

	//Processos em execução
	processData := make([]string, 0, 50)

	//Criação do monitor de temperatura
	tempData := make([]float64, 0, 50)

	//Criação do monitor de uso da cpu
	cpuData := make([]float64, 0, 40)

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

		currentTemp := tempData[len(tempData)-1]
		if currentTemp >= 70 {
			temp.LineColor = ui.ColorRed
			temp.TitleStyle.Fg = ui.ColorRed
			tempStat.BorderStyle.Fg = ui.ColorRed
		} else if currentTemp >= 60 {
			temp.LineColor = ui.ColorYellow
			temp.TitleStyle.Fg = ui.ColorYellow
			tempStat.BorderStyle.Fg = ui.ColorYellow
		} else {
			temp.LineColor = ui.ColorCyan
			temp.TitleStyle.Fg = ui.ColorCyan
			tempStat.BorderStyle.Fg = ui.ColorCyan
		}

		temp.Title = fmt.Sprintf("%.2f", currentTemp)
		temp.Data = tempData

		//Preenchimento do uso da CPU
		percent, _ := services.PercentUsedCpu()
		cpuData = append(cpuData, percent...)
		if len(cpuData) > 40 {
			cpuData = cpuData[len(cpuData)-40:]
		}

		cpuUsage.Data[0] = cpuData

		//Uso da memória
		memoryData, _ := services.InfoMemory()
		currentMemoryPercent := memoryData.UsedPercent
		memoryUsage.Percent = int(currentMemoryPercent)

		if currentMemoryPercent > 95 {
			memoryUsage.BarColor = ui.ColorRed
			memoryUsage.BorderStyle.Fg = ui.ColorRed
			memoryUsage.TitleStyle.Fg = ui.ColorRed
		} else if currentMemoryPercent > 80 {
			memoryUsage.BarColor = ui.ColorYellow
			memoryUsage.BorderStyle.Fg = ui.ColorYellow
			memoryUsage.TitleStyle.Fg = ui.ColorYellow
		} else {
			memoryUsage.BarColor = ui.ColorGreen
			memoryUsage.BorderStyle.Fg = ui.ColorGreen
			memoryUsage.TitleStyle.Fg = ui.ColorGreen
		}

		ui.Clear()
		ui.Render(header, infoHost, process, tempStat, cpuUsage, memoryUsage)
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
