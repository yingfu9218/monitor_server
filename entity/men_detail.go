package entity

import "github.com/shirou/gopsutil/v4/mem"

type MemDetail struct {
	MemInfo         *mem.VirtualMemoryStat `json:"mem_info"`
	MemTotalStr     string                 `json:"mem_total_str"`
	MemAvailableStr string                 `json:"mem_available_str"`
	MemUsedStr      string                 `json:"mem_used_str"`
	MemFreeStr      string                 `json:"mem_free_str"`
}
