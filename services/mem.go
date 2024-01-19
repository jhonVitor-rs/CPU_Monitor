package services

import (
	"fmt"

	mem "github.com/shirou/gopsutil/v3/mem"
)

func InfoMemory() (*mem.VirtualMemoryStat, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Erro ao obter o status da memória")
		return nil, err
	}

	return stat, nil
}
