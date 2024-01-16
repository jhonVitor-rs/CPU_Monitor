package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		fmt.Printf("Erro ao inicializar o termui: %v\n", err)
		return
	}
	defer ui.Close()
}
