package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"kinexon/containerruntime/app/dtos"
	"math"
	"time"
)

func GetCpuUsage(previousCPU, previousSystem uint64, stat *types.Stats) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(stat.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(stat.CPUStats.SystemUsage) - float64(previousSystem)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(GetCpuCount(stat)) * 100.0
	}

	return math.Floor(cpuPercent*100) / 100
}

func GetMemoryUsage(containerStats *types.Stats) float64 {
	value := (float64(containerStats.MemoryStats.Usage) / float64(containerStats.MemoryStats.Limit)) * 100
	return math.Floor(value*100) / 100
}

func BytesToGb(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func BytesToMb(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func GetCpuCount(containerStats *types.Stats) int {
	cpuCount := 0
	if len(containerStats.CPUStats.CPUUsage.PercpuUsage) > 0 {
		cpuCount = len(containerStats.CPUStats.CPUUsage.PercpuUsage)
	} else {
		cpuCount = int(containerStats.CPUStats.OnlineCPUs)
	}
	return cpuCount
}

func GetContainerStats(ctx context.Context, cli *client.Client, containerID string) (<-chan dtos.Stats, error) {
	statsChan := make(chan dtos.Stats)
	var prevStats types.Stats
	go func() {
		defer close(statsChan)
		for {
			data, err := cli.ContainerStats(ctx, containerID, true)
			if err != nil {
				fmt.Println("Error getting container stats:", err)
				return
			}

			defer data.Body.Close()

			decoder := json.NewDecoder(data.Body)
			var containerStats types.Stats
			var containerNet types.StatsJSON

			_err := decoder.Decode(&containerStats)
			decoder.Decode(&containerNet)
			if _err != nil {
				return
			}

			cpuUsage := GetCpuUsage(prevStats.CPUStats.CPUUsage.TotalUsage, prevStats.CPUStats.SystemUsage, &containerStats)
			stats := dtos.Stats{
				TotalCpus:       GetCpuCount(&containerStats),
				CpuUsage:        cpuUsage,
				UsedMemory:      containerStats.MemoryStats.Usage,
				TotalMemory:     containerStats.MemoryStats.Limit,
				MemoryUsage:     GetMemoryUsage(&containerStats),
				UsedMemoryMb:    BytesToMb(containerStats.MemoryStats.Usage),
				TotalMemoryGb:   BytesToGb(containerStats.MemoryStats.Limit),
				NetworkInput:    containerNet.Networks["eth0"].RxBytes,
				NetworkOutput:   containerNet.Networks["eth0"].TxBytes,
				NetworkInputMb:  BytesToMb(containerNet.Networks["eth0"].RxBytes),
				NetworkOutputMb: BytesToMb(containerNet.Networks["eth0"].TxBytes),
			}

			statsChan <- stats // Send stats to the channel
			prevStats = containerStats
			time.Sleep(time.Millisecond * 1000)
		}
	}()
	return statsChan, nil
}
