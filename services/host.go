package services

import (
	"fmt"

	host "github.com/shirou/gopsutil/v3/host"
)

func InfoStat() (*host.InfoStat, error) {
	stat, err := host.Info()
	if err != nil {
		fmt.Printf("Erro ao obter o status do host")
		return nil, err
	}

	return stat, nil
}

func TemperatureStat() ([]host.TemperatureStat, error) {
	stat, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Printf("Erro ao obter o status da temperatura da CPU")
		return nil, err
	}

	return stat, nil
}
