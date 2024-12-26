package entity

import (
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type BaseInfo struct {
	HostInfo        *host.InfoStat         `json:"host_info"`
	CpuInfo         CpuInfo                `json:"cpu_info"`
	CpuUsed         []float64              `json:"cpu_used"`
	MemInfo         *mem.VirtualMemoryStat `json:"mem_info"`
	MemTotalStr     string                 `json:"mem_total_str"`
	MemAvailableStr string                 `json:"mem_available_str"`
	MemUsedStr      string                 `json:"mem_used_str"`
	MemFreeStr      string                 `json:"mem_free_str"`
}
