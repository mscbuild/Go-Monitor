package main

import (
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/net"
)

type Metrics struct {
	Node      string  `json:"node"`
	CPU       float64 `json:"cpu"`
	Memory    float64 `json:"memory"`
	Disk      float64 `json:"disk"`
	Load      float64 `json:"load"`
	NetworkIn  uint64 `json:"network_in"`
	NetworkOut uint64 `json:"network_out"`
	Timestamp int64   `json:"timestamp"`
}

func Collect(node string) Metrics {
	cpuPercent, _ := cpu.Percent(0, false)
	memStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	loadAvg, _ := load.Avg()
	netIO, _ := net.IOCounters(false)

	return Metrics{
		Node:      node,
		CPU:       cpuPercent[0],
		Memory:    memStat.UsedPercent,
		Disk:      diskStat.UsedPercent,
		Load:      loadAvg.Load1,
		NetworkIn:  netIO[0].BytesRecv,
		NetworkOut: netIO[0].BytesSent,
		Timestamp: time.Now().Unix(),
	}
}
