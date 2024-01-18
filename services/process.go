package services

import (
	"fmt"

	process "github.com/shirou/gopsutil/v3/process"
)

func InfoProcess() ([]*process.Process, error) {
	stat, err := process.Processes()
	if err != nil {
		fmt.Printf("Erro ao obter o status dos processos")
		return nil, err
	}

	return stat, err
}
