package services

import (
	"fmt"
	"time"

	cpu "github.com/shirou/gopsutil/v3/cpu"
)

// PercentUsedCpu retorna a porcentagem de uso da CPU
func PercentUsedCpu() ([]float64, error) {
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Printf("Erro ao obter a porcentagem da CPU: %s\n", err)
		return nil, err
	}

	return percent, nil
}
